package security

import (
	"fmt"
	"twimo/backend/core"
	"twimo/backend/encr"
	"twimo/backend/encr/base"
	"twimo/backend/repository"
)

// Password verification of the user's password
func Password(user core.User, password string, r repository.Repository) (bool, error) {

	// Generate salt
	_, salt, err := encr.Salt(user.Name, base.TwimoBase)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// hash password user has put in
	hashedPasswords, err := encr.Hash(password, salt, base.TwimoBase)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// Compare hashed password and the password in the database
	if hashedPasswords == user.Password {
		return true, nil
	}

	return false, nil
}

// IsHashAndUsernameValid checks if the user with the following
// username is in the database and if the hash is valid
func IsHashAndUsernameValid(userData UserData, r repository.Repository) bool {

	// Check if user is in the database
	user, _ := r.GetUserFromName(userData.Username)
	if user.ID == "" {
		return false
	}

	// Check if the key is valid
	key := user.ToChecksum()
	if key == userData.Key {
		return true
	}
	return false
}
