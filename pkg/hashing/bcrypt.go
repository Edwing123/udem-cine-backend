package hashing

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	cost = 12
)

// Hashes the given password using Bcrypt.
func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(password),
		cost,
	)

	return string(hashedPassword)
}

// Verifies if clearText and password match.
func VerifyPassword(clearTextPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(clearTextPassword),
	)

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}

	return true
}
