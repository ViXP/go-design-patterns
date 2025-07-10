package interpreter

type TokenType byte

// TokenType constant enumerables.
const (
	Title TokenType = iota
	Character
	Speak
	Move
	Eat
	Shoot
	Encounter
	End
)

// Token represents the chunk of the abstract syntax tree.
type Token struct {
	Type  TokenType
	Actor string
	Value string
}

// NewToken constructs the new Token.
func NewToken(tokenType TokenType, actor string, value string) *Token {
	return &Token{Type: tokenType, Actor: actor, Value: value}
}
