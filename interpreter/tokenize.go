package interpreter

import (
	"fmt"
	"strings"
	"unicode"
)

// Tokenize implements lexing part of the interpreter.
func Tokenize(fullString string) []Token {
	var tokens []Token = []Token{}

	lines := strings.Split(fullString, ";")

	fmt.Println(lines[len(lines)-1])

	for _, line := range lines {
		var tokenType TokenType
		var operands []string = []string{}

		expressions := strings.Split(line, " ")

		for _, exp := range expressions {
			exp = strings.TrimFunc(exp, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			})

			switch strings.ToLower(exp) {
			case "title":
				tokenType = Title
			case "character":
				tokenType = Character
			case "say":
				tokenType = Speak
			case "move":
				tokenType = Move
			case "eats":
				tokenType = Eat
			case "shoots":
				tokenType = Shoot
			case "encounter":
				tokenType = Encounter
			case "end":
				tokenType = End
			case "":
				tokenType = 255
			default:
				operands = append(operands, exp)
			}
		}

		var value string
		var actor string

		if len(operands) > 1 {
			actor = operands[0]
			value = strings.Join(operands[1:], " ")
		} else if len(operands) == 1 {
			actor = operands[0]
		}
		if tokenType != 255 {
			tokens = append(tokens, *NewToken(tokenType, actor, value))
		}
	}

	return tokens
}
