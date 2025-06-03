package chain_of_responsibility

// AddPatty is the cooking process that adds the patty to the meal.
type AddPatty struct {
	CookingStep
}

// process processes the cooking step.
func (ap *AddPatty) process() {
	ap.Meal.AddLayer("Grilled patty added")
}

// Handle handles the cooking step.
func (ap *AddPatty) Handle() bool {
	return handleStep(ap)
}

// Interface implementation assertion.
var _ Chainable = &AddPatty{}
