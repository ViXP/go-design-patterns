// Package state implements the State design pattern.
// The Context interface defines the common Actioner behavior used to handle actions on a specific subject (in this
// case – the theatre Ticket).
// The State interface shares these common methods across all concrete states. Each specific State knows how to handle
// transitions to other states or can fall back to the default behavior described in the BaseState struct. Every Context
// (Ticket) delegates the execution of its actions to the current State it holds.
package state

func Run() {
	// Creating the new ticket
	ticket := NewTicket()

	// Trying to reserve the ticket:
	ticket.Reserve()
	ticket.GetCurrentState()

	// Trying to buy the ticket:
	ticket.Buy()
	ticket.GetCurrentState()

	// Trying to cancel the purchase:
	ticket.Return()
	ticket.GetCurrentState()

	// Trying to use the returned ticket (ends with error):
	ticket.Use()
}
