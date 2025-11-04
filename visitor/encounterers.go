package visitor

import "fmt"

// Priest represents a holy Encounterer who uses faith and ritual to fight evil. His abilities are most effective
// against Demons and Vampires, less so against Zombies or Witches.
type Priest struct{}

// String returns the stringified name of the Priest.
func (p *Priest) String() string {
	return "The Priest ✝️ "
}

// FaceDemon allows the Priest to interact with the Demon Confronter.
func (p *Priest) FaceDemon(v Confronter) {
	v.Hit(100)
	fmt.Printf("%v says: The power of Christ compels you! And the Demon be gone!\n%v", p, v)
}

// FaceGhost allows the Priest to interact with the Ghost Confronter.
func (p *Priest) FaceGhost(v Confronter) {
	v.Hit(10)
	fmt.Printf("%v says: 'The power of Christ compels you!' But it is not very efficient.\n%v", p, v)
}

// FaceVampire allows the Priest to interact with the Vampire Confronter.
func (p *Priest) FaceVampire(v Confronter) {
	v.Hit(50)
	fmt.Printf("%v shows the cross to the vampire and uses the holy water.\n%v", p, v)
}

// FaceWitch allows the Priest to interact with the Witch Confronter.
func (p *Priest) FaceWitch(v Confronter) {
	v.Hit(0)
	fmt.Printf("%v reads the Bible. The witch is not interested.\n%v", p, v)
}

// FaceZombie allows the Priest to interact with the Zombie Confronter.
func (p *Priest) FaceZombie(v Confronter) {
	v.Hit(1)
	fmt.Printf("%v hits the zombie with his cross. It is not effective.\n%v", p, v)
}

// GhostBuster is a modern hero who relies on gadgets and proton technology. Extremely effective against Ghosts;
// less so against physical foes.
type GhostBuster struct{}

// String returns the stringified name of the GhostBuster.
func (g *GhostBuster) String() string {
	return "The Ghost Buster 🪤 "
}

// FaceDemon allows the GhostBuster to interact with the Demon Confronter.
func (g *GhostBuster) FaceDemon(v Confronter) {
	v.Hit(25)
	fmt.Printf("%v pulls out a proton pack, but it only annoys the Demon.\n%v", g, v)
}

// FaceGhost allows the GhostBuster to interact with the Ghost Confronter.
func (g *GhostBuster) FaceGhost(v Confronter) {
	v.Hit(100)
	fmt.Printf("%v traps the Ghost in a glowing cube.\n%v", g, v)
}

// FaceVampire allows the GhostBuster to interact with the Vampire Confronter.
func (g *GhostBuster) FaceVampire(v Confronter) {
	v.Hit(0)
	fmt.Printf("%v shines a ghost scanner at the Vampire — it does absolutely nothing.\n%v", g, v)
}

// FaceWitch allows the GhostBuster to interact with the Witch Confronter.
func (g *GhostBuster) FaceWitch(v Confronter) {
	v.Hit(10)
	fmt.Printf("%v tries to vacuum up the Witch with his proton pack. She laughs but looks a bit concerned.\n%v", g, v)
}

// FaceZombie allows the GhostBuster to interact with the Zombie Confronter.
func (g *GhostBuster) FaceZombie(v Confronter) {
	v.Hit(30)
	fmt.Printf("%v sprays the Zombie with ectoplasmic goo. Surprisingly effective!\n%v", g, v)
}

// Soldier represents a disciplined fighter with raw firepower and tactical efficiency. Best suited for physical
// threats; struggles against ethereal beings.
type Soldier struct{}

// String returns the stringified name of the Soldier.
func (s *Soldier) String() string {
	return "The Soldier 🪖 "
}

// FaceDemon allows the Soldier to interact with the Demon Confronter.
func (s *Soldier) FaceDemon(v Confronter) {
	v.Hit(40)
	fmt.Printf("%v unloads a full magazine into the Demon. It slows down... slightly.\n%v", s, v)
}

// FaceGhost allows the Soldier to interact with the Ghost Confronter.
func (s *Soldier) FaceGhost(v Confronter) {
	v.Hit(0)
	fmt.Printf("%v shoots right through the Ghost. He now questions his life choices.\n%v", s, v)
}

// FaceVampire allows the Soldier to interact with the Vampire Confronter.
func (s *Soldier) FaceVampire(v Confronter) {
	v.Hit(70)
	fmt.Printf("%v fires silver bullets blessed by the chaplain. Bullseye!\n%v", s, v)
}

// FaceWitch allows the Soldier to interact with the Witch Confronter.
func (s *Soldier) FaceWitch(v Confronter) {
	v.Hit(100)
	fmt.Printf("%v headshots the Witch without even blinking. Rude.\n%v", s, v)
}

// FaceZombie allows the Soldier to interact with the Zombie Confronter.
func (s *Soldier) FaceZombie(v Confronter) {
	v.Hit(100)
	fmt.Printf("%v headshots the Zombie without even blinking. Classic!\n%v", s, v)
}

// VampireHunter specializes in exterminating vampires, but takes on other foes with variable success.
type VampireHunter struct{}

// String returns the stringified name of the VampireHunter.
func (v *VampireHunter) String() string {
	return "The Vampire Hunter 🗡️ "
}

// FaceDemon allows the VampireHunter to interact with the Demon Confronter.
func (vh *VampireHunter) FaceDemon(v Confronter) {
	v.Hit(30)
	fmt.Printf("%v waves his silver-tipped stake at the Demon. It’s mildly impressed.\n%v", vh, v)
}

// FaceGhost allows the VampireHunter to interact with the Ghost Confronter.
func (vh *VampireHunter) FaceGhost(v Confronter) {
	v.Hit(5)
	fmt.Printf("%v swings his holy dagger at the Ghost, cutting nothing but air.\n%v", vh, v)
}

// FaceVampire allows the VampireHunter to interact with the Vampire Confronter.
func (vh *VampireHunter) FaceVampire(v Confronter) {
	v.Hit(100)
	fmt.Printf("%v drives a wooden stake through the Vampire’s heart and mutters a prayer.\n%v", vh, v)
}

// FaceWitch allows the VampireHunter to interact with the Witch Confronter.
func (vh *VampireHunter) FaceWitch(v Confronter) {
	v.Hit(10)
	fmt.Printf("%v throws blessed garlic at the Witch. She’s offended, but also sneezes.\n%v", vh, v)
}

// FaceZombie allows the VampireHunter to interact with the Zombie Confronter.
func (vh *VampireHunter) FaceZombie(v Confronter) {
	v.Hit(5)
	fmt.Printf("%v pokes the Zombie with a silver cross, but it only smells bad.\n%v", vh, v)
}

// Interface implementation assertions.
var (
	_ Encounterer = &Priest{}
	_ Encounterer = &GhostBuster{}
	_ Encounterer = &Soldier{}
	_ Encounterer = &VampireHunter{}
)
