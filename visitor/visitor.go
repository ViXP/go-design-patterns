// Package visitor implements the Visitor design pattern.
// The Encounterer interface defines the common interface for all visitors (heroes) that can interact with various
// elements (monsters) that implement Confronter interface.
// The double dispatch mechanism is achieved through the Confront method, allowing each hero and monster combination to
// define unique interaction logic. In this implementation, different types of heroes (Encounterers) enter a haunted
// house and confront various supernatural creatures (Confronters).
package visitor

func Run() {
	hh := NewHauntedHouse()
	party := NewParty()

	// Add enemies to the Haunted House:
	hh.Add(&Witch{80}).Add(&Demon{100}).Add(&Ghost{100}).Add(&Ghost{100}).Add(&Zombie{100}).Add(&Demon{50})
	hh.Add(&Vampire{100})

	// Add heroes to the Party:
	party.Add(&VampireHunter{}).Add(&GhostBuster{}).Add(&Priest{}).Add(&Soldier{}).Add(&Soldier{})

	// Let's fight!
	hh.Visit(party)
}
