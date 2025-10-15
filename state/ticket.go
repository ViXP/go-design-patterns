package state

import "fmt"

// Ticket is the specific Context that is able to execute actions, by delegating them to the current internal State.
type Ticket struct {
	currentState Actioner
}

// Buy delegates Buy operation to the current State.
func (t *Ticket) Buy() {
	t.currentState.Buy()
}

// CheckIn delegates CheckIn operation to the current State.
func (t *Ticket) CheckIn() {
	t.currentState.CheckIn()
}

// Reserve delegates Reserve operation to the current State.
func (t *Ticket) Reserve() {
	t.currentState.Reserve()
}

// Return delegates Return operation to the current State.
func (t *Ticket) Return() {
	t.currentState.Return()
}

// Use delegates Use operation to the current State.
func (t *Ticket) Use() {
	t.currentState.Use()
}

// ChangeState updates the current State of the Ticket.
func (t *Ticket) ChangeState(state Actioner) {
	t.currentState = state
}

// GetCurrentState prints and returns the current State of the Ticket.
func (t *Ticket) GetCurrentState() Actioner {
	fmt.Println("Current state of the ticket:", t.currentState)
	return t.currentState
}

// NewTicket creates the new Ticket in it's default State (Available).
func NewTicket() *Ticket {
	ticket := Ticket{}
	NewAvailableState(&ticket)
	ticket.GetCurrentState()
	return &ticket
}

// Interface implementation assertion.
var _ Context = &Ticket{}
