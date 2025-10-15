package state

import "fmt"

// ReservedState represents the "booked" mode for the Context. However, it is still can be returned, as well as bought.
type ReservedState struct {
	BaseState
}

// Buy transitions the Context from Reserved to Paid.
func (rs *ReservedState) Buy() {
	NewPaidState(rs.subject)
	fmt.Println("The ticket has been successfully paid for!")
}

// Return transitions the Context from Reserved to Cancelled.
func (rs *ReservedState) Return() {
	NewCancelledState(rs.subject)
	fmt.Println("The reservation has been cancelled.")
}

// String returns the human-readable name of the current State.
func (rs *ReservedState) String() string {
	return "Reserved"
}

// NewReservedState constructs the new ReservedState and assigns it as the current State of the provided Context.
func NewReservedState(subject Context) *ReservedState {
	state := ReservedState{BaseState{subject}}
	subject.ChangeState(&state)
	return &state
}

// Interface implementation assertion.
var _ State = &ReservedState{}
