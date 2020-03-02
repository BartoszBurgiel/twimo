package core

import (
	"strings"
	"twimo/backend/encr"
	"twimo/backend/encr/base"
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

	u.Email, _ = key.Lock([]byte(u.Email), base.TwimoBase)
}

// Decode user's personal data with kihmo
func (u *User) Decode() {
	// Calculate the key
	key := u.generateUserKey()

	// Encode name password and email
	email, _ := key.Unlock(u.Email, base.TwimoBase)
	u.Email = string(email)
}

// Generate code from user's id
func (u User) generateUserKey() encr.Key {

	// KeyCode = last bit from comment code
	dividedID := strings.SplitAfter(u.ID, "-")
	code := dividedID[len(dividedID)-1]
	key, _ := encr.StringToKey(code)
	return key
}

// ToChecksum returns a hash of all user's data in TwimoBase
func (u User) ToChecksum() (sum string) {

	// Decode information
	u.Decode()

	// Gather all user's information
	sum += u.Name
	sum += u.ID
	sum += u.Email
	u.Encode()

	// Generate salt
	_, salt, _ := encr.Salt(u.Name, base.TwimoBase)

	// Hash and replace salt
	sum, _ = encr.Hash(sum, salt, base.TwimoBase)

	return sum
}
