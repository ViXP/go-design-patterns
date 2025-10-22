// Package strategy implements the Strategy design pattern.
// The User struct acts as a Context that can change its behavior dynamically at runtime by switching between different
// UserDataEncoder strategies. Each concrete strategy (MD5HashingStrategy, Base64EncodingStrategy or
// ReverseEncodingStrategy) defines a unique algorithm for encoding user data such as login, name, and surname.
package strategy

import "fmt"

func Run() {
	// Initialize the User with default encoding strategy (MD5 hashing):
	user := NewUser("crypticUser", "John", "Wayne", '_')
	fmt.Println("MD5 hashed User:", user)

	// Change the encoding strategy to Base64 (reversable):
	user.ChangeStrategy(Base64)
	fmt.Println("Base64 encoded User:", user)

	// Change the encoding strategy to Reversed custom strategy:
	user.ChangeStrategy(Reversed)
	fmt.Println("Reversed encoded User:", user)
}
