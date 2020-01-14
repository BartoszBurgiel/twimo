package dummygenerator

import (
	"kihmo"
	"strings"
)

// Stores all data of the user needed to compose the querries
type user struct {
	name        string
	password    string
	email       string
	comments    string
	favlocation string
	ratings     string
	id          string
}

// Stores all data of the location needed to compose the querries
type location struct {
	name     string
	coordX   int
	coordY   int
	desc     string
	comments string
	users    string
	ratings  string
	id       string
}

/*
BARTOSZ'S PART
*/

// Encode user's personal information with kihmo
func (u *user) Encode() {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.comments, "-")
	code := dividedComments[len(dividedComments)-1]

	// Generate key
	key, _ := kihmo.StringToKey(code)

	// Encode
	u.name, _ = key.Lock([]byte(u.name), 60)
	u.password, _ = key.Lock([]byte(u.password), 60)
	u.email, _ = key.Lock([]byte(u.email), 60)
}

// Decode user's personal information with kihmo
func (u *user) Decode() {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.comments, "-")
	code := dividedComments[len(dividedComments)-1]

	// Generate key
	key, _ := kihmo.StringToKey(code)

	// Encode
	name, _ := key.Unlock(u.name, 60)
	password, _ := key.Unlock(u.password, 60)
	email, _ := key.Unlock(u.email, 60)

	u.name = string(name)
	u.password = string(password)
	u.email = string(email)
}
