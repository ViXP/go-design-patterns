package interpreter

import (
	"fmt"
	"strings"
)

// Parse implements Token parsing functionality to the string.
func Parse(tokens []Token) string {
	var parsed strings.Builder = strings.Builder{}

	for _, token := range tokens {
		var line string
		switch token.Type {
		case Title:
			line = fmt.Sprintf("Here is the story of \"%s %s\":\n", token.Actor, token.Value)
		case Character:
			line = fmt.Sprintf("Now let us meet our %s!\n", token.Actor)
		case Speak:
			line = fmt.Sprintf("- \"%s\" – said %s.\n", token.Value, token.Actor)
		case Move:
			line = fmt.Sprintf("Then %s moved to %s.\n", token.Actor, token.Value)
		case Eat:
			line = fmt.Sprintf("But %s ate %s :-(\n", token.Actor, token.Value)
		case Shoot:
			line = fmt.Sprintf("%s shot %s!!!\n", token.Actor, token.Value)
		case Encounter:
			line = fmt.Sprintf("Suddenly, %s saw %s!\n", token.Actor, token.Value)
		case End:
			line = "This is how our story ends..."
		default:
			panic("unknown token! " + string(token.Type))
		}
		parsed.WriteString(line)
	}

	return parsed.String()
}
