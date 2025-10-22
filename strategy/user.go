package strategy

import (
	"fmt"
	"strings"
)

// User represents a the user entity whose data can be encoded using different strategies.
type User struct {
	Login              string
	Name               string
	Surname            string
	separator          rune
	encryptionStrategy UserDataEncoder
}

// ChangeStrategy switches the encoding algorithm used for user data.
func (u *User) ChangeStrategy(strategy Strategy) {
	var encryptionStrategy UserDataEncoder
	switch strategy {
	default:
		encryptionStrategy = &MD5HashingStrategy{}
	case Base64:
		encryptionStrategy = &Base64EncodingStrategy{}
	case Reversed:
		encryptionStrategy = &ReverseEncodingStrategy{}
	}
	u.encryptionStrategy = encryptionStrategy
}

// String returns stringified version of encoded user data, with the parts separated by predefined delimiter.
func (u *User) String() string {
	var buffer strings.Builder = strings.Builder{}

	buffer.WriteString(u.encryptionStrategy.EncodeLogin(u))
	buffer.WriteRune(u.separator)
	buffer.WriteString(u.encryptionStrategy.EncodeName(u))
	buffer.WriteRune(u.separator)
	buffer.WriteString(u.encryptionStrategy.EncodeSurname(u))
	return buffer.String()
}

// NewUser constructs a user with assigned parameters and default encoder
func NewUser(login, name, surname string, separator rune) *User {
	user := &User{Login: login, Name: name, Surname: surname, separator: separator}
	user.ChangeStrategy(MD5)
	return user
}

// Interfaces implementation assertion.
var (
	_ fmt.Stringer    = &User{}
	_ StrategyChanger = &User{}
)
