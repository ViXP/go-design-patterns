package chain_of_responsibility

import (
	"fmt"
	"math/rand"
)

// AddSauce is the cooking process that adds one of the random sauce to the meal.
type AddSauce struct {
	CookingStep
}

// Handle handles the cooking step.
func (a *AddSauce) Handle() bool {
	return handleStep(a)
}

// process processes the cooking step.
func (a *AddSauce) process() {
	sauces := []string{"Ketchup", "Mayonnaise", "Mustard", "BBQ", "Tartar", "Salsa", "Garlic"}

	a.Meal.AddLayer(fmt.Sprintf("Adding the %s sauce", sauces[rand.Intn(len(sauces))]))
}

// Interface implementation assertion.
var _ Chainable = &AddSauce{}
