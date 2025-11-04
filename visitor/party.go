package visitor

import (
	"fmt"
	"strings"
)

// Party is the composite for the Encounterers party, with the ability of adding to party and printing the information
// about it.
type Party struct {
	Heroes []Encounterer
}

// String returns a formatted string describing all heroes in the Party.
func (p *Party) String() string {
	sb := strings.Builder{}
	for _, h := range p.Heroes {
		sb.WriteString(fmt.Sprintf("- %v\n", h))
	}
	return sb.String()
}

// Add allows to add the Encounterers to the Party.
func (p *Party) Add(hero any) Extendable {
	if h, ok := hero.(Encounterer); ok {
		p.Heroes = append(p.Heroes, h)
	}

	return p
}

// NewParty creates the new Party of Encounterers.
func NewParty() *Party {
	return &Party{[]Encounterer{}}
}

// Interface implementation assertions.
var (
	_ Extendable   = &Party{}
	_ fmt.Stringer = &Party{}
)
