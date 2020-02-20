package core

import (
	"kihmo"
	"kihmo/base"
	"strings"
)

// User struct represents a user
type User struct {
	Name        string
	Password    string
	Email       string
	FavLocation string
	ID          string
}

// Encode user's personal data with kihmo
func (u *User) Encode() {
	// Calculate the key
	key := u.generateUserKey()

	u.Email, _ = key.Lock([]byte(u.Email), base.Base85)
}

// Decode user's personal data with kihmo
func (u *User) Decode() {
	// Calculate the key
	key := u.generateUserKey()

	// Encode name password and email
	email, _ := key.Unlock(u.Email, base.Base85)
	u.Email = string(email)
}

// Generate code from user's id
func (u User) generateUserKey() kihmo.Key {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.FavLocation, "-")
	code := dividedComments[len(dividedComments)-1]
	key, _ := kihmo.StringToKey(code)
	return key
}

// ToChecksum returns a hash of all user's data in base85
func (u User) ToChecksum() (sum string) {
	return sum
}
