package security

import (
	"fmt"
	"kihmo"
	"twimo/backend/repository"
)

// VerifyPassword the user with given Id has put in
func VerifyPassword(userID, password string, r repository.Repository) (bool, error) {

	// Get user data
	user, err := r.GetUser(userID)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// Generate salt
	_, salt, err := kihmo.Salt(user.Name, 60)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// hash password
	hashedPasswords, err := kihmo.Hash(userID, salt, 60)
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
