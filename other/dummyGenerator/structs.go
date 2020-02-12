package dummygenerator

import (
	"kihmo"
	"kihmo/base"
	"strings"
)

// Stores all data of the user needed to compose the querries
type user struct {
	name        string
	password    string
	email       string
	favlocation string
	id          string
}

// Stores all data of the location needed to compose the querries
type location struct {
	name    string
	coordX  int
	coordY  int
	desc    string
	website string
	id      string
}

// stores all data of all ratings
type rating struct {
	value int
	id    string
}

/*
BARTOSZ'S PART
*/

// Encode user's personal information with kihmo
func (u *user) Encode() {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.favlocation, "-")
	code := dividedComments[len(dividedComments)-1]

	// Generate key
	key, _ := kihmo.StringToKey(code)

	u.email, _ = key.Lock([]byte(u.email), base.Base85)
}

// Decode user's personal information with kihmo
func (u *user) Decode() {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.favlocation, "-")
	code := dividedComments[len(dividedComments)-1]

	// Generate key
	key, _ := kihmo.StringToKey(code)

	email, _ := key.Unlock(u.email, base.Base85)

	u.email = string(email)
}
