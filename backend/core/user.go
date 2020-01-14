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
	u.Name, _ = key.Lock([]byte(u.Name), 60)
	u.Password, _ = key.Lock([]byte(u.Password), 60)
	u.Email, _ = key.Lock([]byte(u.Email), 60)
}

// Decode user's personal data with kihmo
func (u *User) Decode() {
	// Calculate the key
	key := u.generateUserKey()

	// Encode name password and email
	name, _ := key.Unlock(u.Name, 60)
	password, _ := key.Unlock(u.Password, 60)
	email, _ := key.Unlock(u.Email, 60)

	u.Name = string(name)
	u.Password = string(password)
	u.Email = string(email)
}

// Generate code from user's id
func (u User) generateUserKey() kihmo.Key {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.Comments, "-")
	code := dividedComments[len(dividedComments)-1]
	key, _ := kihmo.StringToKey(code)
	return key
}
