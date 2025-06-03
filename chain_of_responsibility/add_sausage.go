package chain_of_responsibility

import "fmt"

// Sausage type enums
const (
	Wiener byte = iota
	Frankfurter
	Bratwurst
	Soy
)

// AddSausage is the cooking process that adds the sausage of the specific type to the meal.
type AddSausage struct {
	sausageType byte
	CookingStep
}

// Handle handles the cooking step.
func (a *AddSausage) Handle() bool {
	return handleStep(a)
}

// process processes the cooking step.
func (a *AddSausage) process() {
	newLayer := fmt.Sprintf("The %s sausage is added", a.getType())
	a.Meal.AddLayer(newLayer)
}

func (a *AddSausage) getType() string {
	switch a.sausageType {
	case Frankfurter:
		return "Frankfurter"
	case Bratwurst:
		return "Bratwurst"
	case Soy:
		return "Soy"
	default:
		return "Wiener"
	}
}

// Interface implementation assertion.
var _ Chainable = &AddSausage{}
