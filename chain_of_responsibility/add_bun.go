package chain_of_responsibility

import "fmt"

// Position placement enums
const (
	bottom byte = iota
	top
)

// AddBun is the cooking process that adds the bun in the specific place of the meal.
type AddBun struct {
	CookingStep
	position byte
}

// process processes the cooking step.
func (ab *AddBun) process() {
	newLayer := fmt.Sprintf("The bun is placed on %s", ab.placement())
	ab.Meal.AddLayer(newLayer)
}

func (ab *AddBun) placement() string {
	switch ab.position {
	case top:
		return "top"
	default:
		return "bottom"
	}
}

// Handle handles the cooking step.
func (ab *AddBun) Handle() bool {
	return handleStep(ab)
}

// Interface implementation assertion.
var _ Chainable = &AddBun{}
