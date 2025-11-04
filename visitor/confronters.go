package visitor

import "fmt"

// Demon represents a powerful infernal creature that thrives on chaos and fire. Highly resistant to physical attacks;
// faith-based assaults are most effective.
type Demon struct {
	life int
}

// IsAlive returns true if the Demon still has life points.
func (d *Demon) IsAlive() bool {
	return d.life > 0
}

// Hit reduces the Demon’s life by the given strength value.
func (d *Demon) Hit(strength int) {
	d.life -= strength
}

// String returns a textual representation of the Demon’s current condition.
func (d *Demon) String() string {
	if d.life < 1 {
		return "The Demon 😈 is dead.\n"
	} else {
		return fmt.Sprintf("The Demon 😈 has %v life points\n", d.life)
	}
}

// Confront allows the Demon to face an Encounterer through double dispatch.
func (d *Demon) Confront(v Encounterer) {
	v.FaceDemon(d)
}

// Ghost represents an ethereal, floating spirit with limited ectoplasmic energy. Immune to bullets and blades, but weak
// to scientific gadgets and faith.
type Ghost struct {
	ectoplasm int
}

// IsAlive returns true if the Ghost still has ectoplasmic energy.
func (g *Ghost) IsAlive() bool {
	return g.ectoplasm > 0
}

// Hit reduces the Ghost’s ectoplasmic energy by the given strength value.
func (g *Ghost) Hit(strength int) {
	g.ectoplasm -= strength
}

// String returns a textual representation of the Ghost’s current state.
func (g *Ghost) String() string {
	if g.ectoplasm < 1 {
		return "The Ghost 👻 vanished.\n"
	} else {
		return fmt.Sprintf("The Ghost 👻 has %v%% of ectoplasm\n", g.ectoplasm)
	}
}

// Confront allows the Ghost to face an Encounterer through double dispatch.
func (g *Ghost) Confront(v Encounterer) {
	v.FaceGhost(g)
}

// Vampire represents an undead predator who feeds on the living. Weak against holy items, sunlight, and anyone
// carrying garlic.
type Vampire struct {
	life int
}

// IsAlive returns true if the Vampire still has life points.
func (v *Vampire) IsAlive() bool {
	return v.life > 0
}

// Hit reduces the Vampire’s life by the given strength value.
func (v *Vampire) Hit(strength int) {
	v.life -= strength
}

// String returns a textual representation of the Vampire’s current condition.
func (v *Vampire) String() string {
	if v.life < 1 {
		return "The Vampire 🧛‍♂️ is turned to ashes.\n"
	} else {
		return fmt.Sprintf("The Vampire 🧛‍♂️ has %v life points\n", v.life)
	}
}

// Confront allows the Vampire to face an Encounterer through double dispatch.
func (v *Vampire) Confront(hv Encounterer) {
	hv.FaceVampire(v)
}

// Witch represents a cunning spellcaster skilled in curses and mischief. Difficult to kill, but vulnerable to precision
// and disbelief.
type Witch struct {
	life int
}

// IsAlive returns true if the Witch still has life points.
func (w *Witch) IsAlive() bool {
	return w.life > 0
}

// Hit reduces the Witch’s life by the given strength value.
func (w *Witch) Hit(strength int) {
	w.life -= strength
}

// String returns a textual representation of the Witch’s current condition.
func (w *Witch) String() string {
	if w.life < 1 {
		return "The Witch 🧙‍♀️ is dead.\n"
	} else {
		return fmt.Sprintf("The Witch 🧙‍♀️ has %v life points\n", w.life)
	}
}

// Confront allows the Witch to face an Encounterer through double dispatch.
func (w *Witch) Confront(v Encounterer) {
	v.FaceWitch(w)
}

// Zombie represents a reanimated corpse driven by hunger and confusion. Easy to kill, but is very interested in brains.
type Zombie struct {
	strength int
}

// IsAlive returns true if the Zombie still has remaining strength.
func (z *Zombie) IsAlive() bool {
	return z.strength > 0
}

// Hit reduces the Zombie’s strength by the given amount.
func (z *Zombie) Hit(strength int) {
	z.strength -= strength
}

// String returns a textual representation of the Zombie’s current condition.
func (z *Zombie) String() string {
	if z.strength < 1 {
		return "The Zombie 🧟‍♂️ is not moving anymore.\n"
	} else {
		return fmt.Sprintf("The Zombie 🧟‍♂️ has %v strength\n", z.strength)
	}
}

// Confront allows the Zombie to face an Encounterer through double dispatch.
func (z *Zombie) Confront(v Encounterer) {
	v.FaceZombie(z)
}

// Interface implementation assertions.
var (
	_ Confronter = &Zombie{}
	_ Confronter = &Witch{}
	_ Confronter = &Vampire{}
	_ Confronter = &Ghost{}
	_ Confronter = &Demon{}
)
