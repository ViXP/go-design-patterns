package state

import "fmt"

// AvailableState represents the fresh new state of the Context. It could be bought or reserved.
type AvailableState struct {
	BaseState
}

// Buy transitions the Context from Available to Paid.
func (as *AvailableState) Buy() {
	NewPaidState(as.subject)
	fmt.Println("The ticket has been successfully purchased!")
}

// Reserve transitions the Context from Available to Reserved.
func (as *AvailableState) Reserve() {
	NewReservedState(as.subject)
	fmt.Println("The ticket has been reserved for you!")
}

// String returns the human-readable name of the current State.
func (as *AvailableState) String() string {
	return "Available"
}

// NewAvailableState constructs the new AvailableState and assigns it as the current State of the provided Context.
func NewAvailableState(subject Context) *AvailableState {
	state := AvailableState{BaseState{subject}}
	subject.ChangeState(&state)
	return &state
}

// Interface implementation assertion.
var _ State = &AvailableState{}
