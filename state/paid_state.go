package state

import "fmt"

// PaidState represents the purchased state of the Context. After this, it can either be checked in (used) or refunded.
type PaidState struct {
	BaseState
}

// CheckIn transitions the Context from Paid to Used.
func (ps *PaidState) CheckIn() {
	NewUsedState(ps.subject)
	fmt.Println("The ticket has been successfully checked in!")
}

// Return transitions the Context from Paid to Cancelled (refund).
func (ps *PaidState) Return() {
	NewCancelledState(ps.subject)
	fmt.Println("The ticket has been successfully refunded!")
}

// String returns the human-readable name of the current State.
func (ps *PaidState) String() string {
	return "Paid"
}

// NewPaidState constructs the new PaidState and assigns it as the current State of the provided Context.
func NewPaidState(subject Context) *PaidState {
	state := PaidState{BaseState{subject}}
	subject.ChangeState(&state)
	return &state
}

// Interface implementation assertion.
var _ State = &PaidState{}
