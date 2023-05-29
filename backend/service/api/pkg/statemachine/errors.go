package statemachine

type InvalidEventError struct {
	Event string
	State string
}

func (e InvalidEventError) Error() string {
	return "event " + e.Event + " inappropriate in current consts " + e.State
}

type UnknownEventError struct {
	Event string
}

func (e UnknownEventError) Error() string {
	return "event " + e.Event + " does not exist"
}

type InTransitionError struct {
	Event string
}

func (e InTransitionError) Error() string {
	return "event " + e.Event + " inappropriate because previous transition did not complete"
}

type NotInTransitionError struct{}

func (e NotInTransitionError) Error() string {
	return "transition inappropriate because no consts change in progress"
}

type NoTransitionError struct {
	Err error
}

func (e NoTransitionError) Error() string {
	if e.Err != nil {
		return "no transition with error: " + e.Err.Error()
	}
	return "no transition"
}

type InternalError struct {
	Err error
}

func (e InternalError) Error() string {
	return "internal error on consts transition error" + e.Err.Error()
}
