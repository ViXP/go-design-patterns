package mediator

// ControlTower is the mediator that implements TrafficControl interface.
type ControlTower struct {
	runwayIsFree  bool
	aircraftQueue []Aircraft
}

// NotifyOperationEnding implements the traffic control feedback for the operation ending attempt from the aircraft.
func (c *ControlTower) NotifyOperationEnding(ac Aircraft) {
	if len(c.aircraftQueue) == 0 {
		ac.ReceiveControlMessage("No active operation to end.")
		return
	}
	if c.aircraftQueue[0] == ac {
		c.aircraftQueue = c.aircraftQueue[1:]
		ac.ReceiveControlMessage("Operation completed successfully.")

		if len(c.aircraftQueue) == 0 {
			c.runwayIsFree = true
		}
	} else {
		ac.ReceiveControlMessage("You can't proceed your operation right now, wait for your turn.")
	}
}

// RequestCraftLanding checks the runway, puts aircraft in the queue and notifies if landing can be executed.
func (c *ControlTower) RequestCraftLanding(ac Aircraft) bool {
	ac.ReceiveControlMessage("Request acknowledged: landing reserved.")
	c.aircraftQueue = append(c.aircraftQueue, ac)

	if c.runwayIsFree {
		c.runwayIsFree = false
		ac.ReceiveControlMessage("Runway free: proceed with landing.")
		return true
	} else {
		ac.ReceiveControlMessage("Runway occupied: please wait.")
		return false
	}
}

// RequestCraftTakeoff checks the runway, puts aircraft in the queue and notifies if taking off can be executed.
func (c *ControlTower) RequestCraftTakeoff(ac Aircraft) bool {
	ac.ReceiveControlMessage("Request acknowledged: take off reserved.")
	c.aircraftQueue = append(c.aircraftQueue, ac)

	if c.runwayIsFree {
		c.runwayIsFree = false
		ac.ReceiveControlMessage("Runway free: proceed with takeoff.")
		return true
	} else {
		ac.ReceiveControlMessage("Runway occupied: please wait.")
		return false
	}
}

// NewControlTower is the ControlTower constructor.
func NewControlTower() *ControlTower {
	return &ControlTower{true, []Aircraft{}}
}

// Interface implementation assertion.
var _ TrafficControl = &ControlTower{}
