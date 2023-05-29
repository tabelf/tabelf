package statemachine

import (
	"context"
	"sync"
)

type Callback func(context.Context, *Event) error

type Events []EventDesc

type Callbacks map[string]Callback

type FSM struct {
	state       string
	transitions map[eventKey]string
	callbacks   map[string]Callback
	transition  func() error
	stateMu     sync.RWMutex
	eventMu     sync.Mutex
}

type eventKey struct {
	// 事件的名称.
	event string

	// 当前事件的起源/起点
	src string
}

type EventDesc struct {
	Name     string
	Src      string
	Dst      string
	Callback Callback
}

func NewFSM(initial string, events []EventDesc) *FSM {
	f := &FSM{
		state:       initial,
		transitions: make(map[eventKey]string),
		callbacks:   make(map[string]Callback),
	}
	for _, e := range events {
		f.transitions[eventKey{e.Name, e.Src}] = e.Dst
		if e.Callback != nil {
			f.callbacks[e.Name] = e.Callback
		}
	}
	return f
}

func (f *FSM) GetState() string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return f.state
}

func (f *FSM) SetState(state string) {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()
	f.state = state
}

func (f *FSM) EqualState(state string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return state == f.state
}

// 如果事件会发生返回 true.
func (f *FSM) Can(event string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	_, ok := f.transitions[eventKey{event, f.state}]
	return ok && (f.transition == nil)
}

// 如果事件不可能发生返回 true.
func (f *FSM) Cannot(event string) bool {
	return !f.Can(event)
}

// 返回当前状态下可以的转换的列表.
func (f *FSM) AvailableTransitions() []string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	var transitions []string
	for key := range f.transitions {
		if key.src == f.state {
			transitions = append(transitions, key.event)
		}
	}
	return transitions
}

func (f *FSM) Event(ctx context.Context, event string, args map[string]interface{}) error {
	f.eventMu.Lock()
	defer f.eventMu.Unlock()

	f.stateMu.RLock()
	defer f.stateMu.RUnlock()

	if f.transition != nil {
		return InTransitionError{event}
	}
	dst, ok := f.transitions[eventKey{event, f.state}]
	if !ok {
		for t := range f.transitions {
			if t.event == event {
				return InvalidEventError{event, f.state}
			}
		}
		return UnknownEventError{event}
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	e := &Event{
		FSM:        f,
		Event:      event,
		Src:        f.state,
		Dst:        dst,
		Args:       args,
		canceled:   false,
		cancelFunc: cancel,
	}
	if f.state == dst {
		if err := f.EventCallbacks(ctx, e); err != nil {
			return err
		}
		return NoTransitionError{nil}
	}
	f.transition = func() (err error) {
		err = f.EventCallbacks(ctx, e)
		if err != nil {
			return err
		}
		f.stateMu.Lock()
		f.state = dst
		f.stateMu.Unlock()
		return nil
	}
	// 提前释放，在执行回调的时候，方法里面加了写锁。否则会阻塞.
	f.stateMu.RUnlock()
	// 为了解决上面的 defer 加的释放锁问题.
	defer f.stateMu.RLock()
	err := f.doTransition()
	if err != nil {
		return InternalError{err}
	}
	return nil
}

func (f *FSM) Transition() error {
	f.eventMu.Lock()
	defer f.eventMu.Unlock()
	return f.doTransition()
}

func (f *FSM) doTransition() error {
	if f.transition == nil {
		return NotInTransitionError{}
	}
	defer func() {
		f.transition = nil
	}()
	return f.transition()
}

func (f *FSM) EventCallbacks(ctx context.Context, e *Event) (err error) {
	if fn, ok := f.callbacks[e.Event]; ok {
		return fn(ctx, e)
	}
	return nil
}
