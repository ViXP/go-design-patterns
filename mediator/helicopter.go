package mediator

import "fmt"

// Helicopter is a colleague that implements Airctaft interface.
type Helicopter struct {
	ID         string
	controller TrafficControl
}

// CompleteOperation tries to finish the previously requested operation.
func (h *Helicopter) CompleteOperation() {
	h.controller.NotifyOperationEnding(h)
}

// ReceiveControlMessage logs the individual feedback from the mediator.
func (h *Helicopter) ReceiveControlMessage(message string) {
	fmt.Printf("[%v received control message]: %v\n", h.ID, message)
}

// RequestLanding notifies mediator on the landing necessity.
func (h *Helicopter) RequestLanding() bool {
	return h.controller.RequestCraftLanding(h)
}

// RequestTakeOff notifies mediator on the taking off necessity.
func (h *Helicopter) RequestTakeOff() bool {
	return h.controller.RequestCraftTakeoff(h)
}

// NewHelicopter is the Helicopter constructor, that receives the identifier and traffic control mediator as params.
func NewHelicopter(id string, ct TrafficControl) *Helicopter {
	return &Helicopter{id, ct}
}

// Interface implementation assertion.
var _ Aircraft = &Helicopter{}
