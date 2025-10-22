package strategy

import "strings"

// ReverseEncodingStrategy implements the custom encoder that encodes the User's data by reversing the elements of it.
type ReverseEncodingStrategy struct{}

// reverse returns the reversed form of the given string.
func (r *ReverseEncodingStrategy) reverse(original string) string {
	reversed := strings.Builder{}
	originalBytes := []byte(original)
	counter := len(originalBytes) - 1

	for {
		reversed.WriteByte(originalBytes[counter])
		counter--
		if counter < 0 {
			break
		}
	}
	return reversed.String()
}

// EncodeLogin allows to encode the User's login.
func (r *ReverseEncodingStrategy) EncodeLogin(user *User) string {
	return r.reverse(user.Login)
}

// EncodeName allows to encode the User's name.
func (r *ReverseEncodingStrategy) EncodeName(user *User) string {
	return r.reverse(user.Name)
}

// EncodeSurname allows to encode the User's surname.
func (r *ReverseEncodingStrategy) EncodeSurname(user *User) string {
	return r.reverse(user.Surname)
}

// Interface implementation assertion.
var _ UserDataEncoder = &ReverseEncodingStrategy{}
