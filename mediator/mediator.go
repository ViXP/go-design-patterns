// Package mediator implements the Mediator design pattern.
// The Aircraft interface represents the colleague entities that can request landing or takeoff operations, always via
// the mediator.
// The TrafficControl interface represents the mediator that coordinates aircraft operations. It ensures aircraft
// interact only indirectly, through centralized traffic organization.
// ControlTower is the concrete mediator, while Plane and Helicopter are concrete colleagues.
package mediator

func Run() {
	tower := NewControlTower()

	// Construct the specific aircrafts
	plane := NewPlane("Boeing 747", tower)
	helicopter := NewHelicopter("Airbus H225", tower)

	// Request different operations
	plane.RequestLanding()
	helicopter.RequestTakeOff()

	// Try to complete the operation ignoring the queue
	helicopter.CompleteOperation()

	// Try to complete operations in the correct queue order
	plane.CompleteOperation()
	helicopter.CompleteOperation()
}
