package core

import (
	"kihmo"
	"strings"
)

// User struct represents a user
type User struct {
	Name        string
	Password    string
	Email       string
	Comments    string
	FavLocation string
	Ratings     string
	ID          string
}

// Encode user's personal data with kihmo
func (u *User) Encode() {
	// Calculate the key
	key := u.generateUserKey()

	// Encode name password and email
	u.Name = key.Lock([]byte(u.Name))
	u.Password = key.Lock([]byte(u.Password))
	u.Email = key.Lock([]byte(u.Email))
}

// Decode user's personal data with kihmo
func (u *User) Decode() {
	// Calculate the key
	key := u.generateUserKey()

	// Encode name password and email
	u.Name = string(key.Unlock(u.Name))
	u.Password = string(key.Unlock(u.Password))
	u.Email = string(key.Unlock(u.Email))
}

// Generate code from user's id
func (u User) generateUserKey() kihmo.Key {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.Comments, "-")
	code := dividedComments[len(dividedComments)-1]

	return kihmo.StringToKey(code)
}
