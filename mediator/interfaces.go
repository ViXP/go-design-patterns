package mediator

// Aircraft is the common interface for all colleagues.
type Aircraft interface {
	RequestTakeOff() bool
	RequestLanding() bool
	CompleteOperation()
	ReceiveControlMessage(string)
}

// TrafficControl is the common interface for the mediators.
type TrafficControl interface {
	RequestCraftLanding(Aircraft) bool
	RequestCraftTakeoff(Aircraft) bool
	NotifyOperationEnding(Aircraft)
}
