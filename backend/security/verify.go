package security

import (
	"fmt"
	"kihmo"
	"kihmo/base"
	"twimo/backend/core"
	"twimo/backend/repository"
)

// Password verification of the user's password
func Password(user core.User, password string, r repository.Repository) (bool, error) {

	// Generate salt
	_, salt, err := kihmo.Salt(user.Name, base.Base85)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// hash password user has put in
	hashedPasswords, err := kihmo.Hash(password, salt, base.Base85)
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
