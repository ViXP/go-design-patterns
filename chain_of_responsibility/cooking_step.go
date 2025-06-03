package chain_of_responsibility

// CookingStep is the abstract struct that implements the Addable interface and encapsulates the reference to the
// meal that is cooking and the next step in chain.
type CookingStep struct {
	Meal Combinable
	next Chainable
}

// AddNext allows to add the next chainable abstraction to the chain.
func (c *CookingStep) AddNext(nextStep Chainable) Chainable {
	c.next = nextStep
	return c.next
}

// getNext retrieves the next chainable abstraction from the chain.
func (c *CookingStep) getNext() Chainable {
	return c.next
}

// Interface implementation assertion.
var _ Nestable = &CookingStep{}
