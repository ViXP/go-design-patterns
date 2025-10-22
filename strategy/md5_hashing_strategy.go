package strategy

import (
	"crypto/md5"
	"fmt"
)

// MD5HashingStrategy implements the MD5 hashing strategy for the User's data.
type MD5HashingStrategy struct{}

// EncodeLogin allows to encode the User's login.
func (m *MD5HashingStrategy) EncodeLogin(user *User) string {
	bytes := []byte(user.Login)
	return fmt.Sprintf("%x", md5.Sum(bytes))
}

// EncodeName allows to encode the User's name.
func (m *MD5HashingStrategy) EncodeName(user *User) string {
	bytes := []byte(user.Name)
	return fmt.Sprintf("%x", md5.Sum(bytes))
}

// EncodeSurname allows to encode the User's surname.
func (m *MD5HashingStrategy) EncodeSurname(user *User) string {
	bytes := []byte(user.Surname)
	return fmt.Sprintf("%x", md5.Sum(bytes))
}

// Interface implementation assertion.
var _ UserDataEncoder = &MD5HashingStrategy{}
