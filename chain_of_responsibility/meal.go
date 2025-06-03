package chain_of_responsibility

import (
	"fmt"
	"strings"
)

// Meal represents the meal that is cooked through the chain of responsibility.
type Meal struct {
	title  string
	layers []string
}

// AddLayer adds the layer to the meal.
func (m *Meal) AddLayer(layer string) {
	m.layers = append(m.layers, layer)
}

// String returns human readable form of the cooked meal.
func (m *Meal) String() string {
	return fmt.Sprintf("Cooking the %s:\n%s\n", strings.ToUpper(m.title), m.combineLayers())
}

func (m *Meal) combineLayers() string {
	combined := strings.Builder{}

	for _, layer := range m.layers {
		combined.WriteString(layer + "\n")
	}
	return combined.String()
}

// NewMeal is the constructor of the new meal that will be cooked.
func NewMeal(title string) *Meal {
	return &Meal{title, []string{}}
}

// Interface implementation assertion.
var _ Combinable = &Meal{}
