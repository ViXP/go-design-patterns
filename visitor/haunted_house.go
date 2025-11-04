package visitor

import "fmt"

// HauntedHouse represents a collection of Confronters—monsters that can be encountered by a Party of Encounters.
// It provides methods for adding Confronters, initiating encounters, and displaying the results of battles.
type HauntedHouse struct {
	enemies []Confronter
}

// Add allows to add the Confronter in the House.
func (h *HauntedHouse) Add(enemy any) Extendable {
	if e, ok := enemy.(Confronter); ok {
		h.enemies = append(h.enemies, e)
	}

	return h
}

// Visit allows a Party of Encounterers to confront all Confronters within the HauntedHouse. Each hero sequentially
// faces the enemies, and the final battle results are displayed afterward.
func (h *HauntedHouse) Visit(party *Party) {
	fmt.Println("===== The Haunted House contains =====")
	h.describe()

	fmt.Println("===== The Party enters =====")
	fmt.Println(party)

	for num, enemy := range h.enemies {
		fmt.Printf("===== Encounter %d =====\n", num+1)
		for _, hero := range party.Heroes {
			if !enemy.IsAlive() {
				break
			}
			enemy.Confront(hero)
		}
		fmt.Println()
	}

	fmt.Println("===== Battle Results =====")
	h.describe()
}

func (h *HauntedHouse) describe() {
	for _, enemy := range h.enemies {
		fmt.Print(enemy)
	}
	fmt.Println()
}

// NewHauntedHouse creates the new Haunted House with Confronters.
func NewHauntedHouse() *HauntedHouse {
	return &HauntedHouse{[]Confronter{}}
}

// Interface implementation assertion.
var _ Extendable = &HauntedHouse{}
