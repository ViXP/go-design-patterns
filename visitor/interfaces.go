package visitor

import (
	"fmt"
)

// Encounterer defines the common interface for all visitors that can interact with different Confronters
// using double dispatch.
type Encounterer interface {
	FaceGhost(v Confronter)
	FaceDemon(v Confronter)
	FaceVampire(v Confronter)
	FaceZombie(v Confronter)
	FaceWitch(v Confronter)
	fmt.Stringer
}

// Confronter defines the common interface for all elements that can be confronted by Encounterers (visitors).
// It provides methods for engaging in combat, taking damage, reporting their state, and checking whether they are still
// alive.
type Confronter interface {
	Confront(v Encounterer)
	Hit(strength int)
	IsAlive() bool
	fmt.Stringer
}

// Extendable defines the ability to add new participants into a composite structure (Encounterers and Confronters).
type Extendable interface {
	Add(participant any) Extendable
}
