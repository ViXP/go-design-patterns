package strategy

type Strategy byte

// Strategies enum:
const (
	MD5 Strategy = iota
	Base64
	Reversed
)

// UserDataEncoder defines the general interface for all User's data encoding strategies.
type UserDataEncoder interface {
	EncodeName(user *User) string
	EncodeSurname(user *User) string
	EncodeLogin(user *User) string
}

// StrategyChanger defines the changing strategy interface for User's data structures.
type StrategyChanger interface {
	ChangeStrategy(strategy Strategy)
}
