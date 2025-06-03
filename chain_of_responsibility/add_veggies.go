package chain_of_responsibility

import (
	"fmt"
	"math/rand"
	"strings"
)

// AddVeggies is the cooking process that adds 2 random veggies to the meal.
type AddVeggies struct {
	CookingStep
}

// process processes the cooking step.
func (av *AddVeggies) process() {
	veggies := av.randomVeggies()
	av.Meal.AddLayer(fmt.Sprintf("%s added", veggies))
}

// Handle handles the cooking step.
func (av *AddVeggies) Handle() bool {
	return handleStep(av)
}

func (av *AddVeggies) randomVeggies() string {
	result := strings.Builder{}
	veggies := []string{"Tomato", "Lettuce", "Onion", "Pickle", "Olive"}

	for range 2 {
		result.WriteString(fmt.Sprintf("%s, ", veggies[rand.Intn(len(veggies))]))
	}
	return result.String()[:len(result.String())-2]
}

// Interface implementation assertion.
var _ Chainable = &AddVeggies{}
