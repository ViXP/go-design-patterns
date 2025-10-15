package state

// CancelledState represents the liminal State, that could be considered as an alias of the AvailableState.
type CancelledState struct {
	AvailableState
}

// String returns the human-readable name of the current State.
func (cs *CancelledState) String() string {
	return "Cancelled (refund)"
}

// NewCancelledState constructs the new CancelledState and assigns it as the current State of the provided Context.
func NewCancelledState(subject Context) *CancelledState {
	state := CancelledState{AvailableState{BaseState{subject}}}
	subject.ChangeState(&state)
	return &state
}

// Interface implementation assertion.
var _ State = &CancelledState{}
