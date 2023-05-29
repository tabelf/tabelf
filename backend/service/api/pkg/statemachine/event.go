package statemachine

type Event struct {
	FSM        *FSM
	Event      string
	Src        string
	Dst        string
	Args       map[string]interface{}
	canceled   bool
	cancelFunc func()
}
