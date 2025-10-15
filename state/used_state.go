package state

// UsedState represents the final irreversable State of the Context.
type UsedState struct {
	BaseState
}

// String returns the human-readable name of the current State.
func (us *UsedState) String() string {
	return "Used"
}

// NewUsedState constructs the new UsedState and assigns it as the current State of the provided Context.
func NewUsedState(subject Context) *UsedState {
	state := UsedState{BaseState{subject}}
	subject.ChangeState(&state)
	return &state
}

// Interface implementation assertion.
var _ State = &UsedState{}
