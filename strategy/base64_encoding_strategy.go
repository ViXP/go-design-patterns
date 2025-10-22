package strategy

import "encoding/base64"

// Base64EncodingStrategy implements the Base64 encoding strategy for the User's data.
type Base64EncodingStrategy struct{}

// EncodeLogin allows to encode the User's login.
func (b *Base64EncodingStrategy) EncodeLogin(user *User) string {
	return base64.StdEncoding.EncodeToString([]byte(user.Login))
}

// EncodeName allows to encode the User's name.
func (b *Base64EncodingStrategy) EncodeName(user *User) string {
	return base64.StdEncoding.EncodeToString([]byte(user.Name))
}

// EncodeSurname allows to encode the User's surname.
func (b *Base64EncodingStrategy) EncodeSurname(user *User) string {
	return base64.StdEncoding.EncodeToString([]byte(user.Surname))
}

// Interface implementation assertion.
var _ UserDataEncoder = &Base64EncodingStrategy{}
