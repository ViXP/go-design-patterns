package mediator

import "fmt"

// Plane is a colleague that implements Airctaft interface.
type Plane struct {
	ID         string
	controller TrafficControl
}

// CompleteOperation tries to finish the previously requested operation.
func (p *Plane) CompleteOperation() {
	p.controller.NotifyOperationEnding(p)
}

// ReceiveControlMessage logs the individual feedback from the mediator.
func (p *Plane) ReceiveControlMessage(message string) {
	fmt.Printf("[%v received control message]: %v\n", p.ID, message)
}

// RequestLanding notifies mediator on the landing necessity.
func (p *Plane) RequestLanding() bool {
	return p.controller.RequestCraftLanding(p)
}

// RequestTakeOff notifies mediator on the taking off necessity.
func (p *Plane) RequestTakeOff() bool {
	return p.controller.RequestCraftTakeoff(p)
}

// NewPlane is the Plane constructor, that receives the identifier and traffic control mediator as params.
func NewPlane(id string, ct TrafficControl) *Plane {
	return &Plane{id, ct}
}

// Interface implementation assertion.
var _ Aircraft = &Plane{}
