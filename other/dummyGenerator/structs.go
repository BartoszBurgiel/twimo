package dummygenerator

import (
	"strings"
	"twimo/backend/encr"
	"twimo/backend/encr/base"
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
	price   int
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

// Encode user's personal information with encr
func (u *user) Encode() {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.id, "-")
	code := dividedComments[len(dividedComments)-1]

	// Generate key
	key, _ := encr.StringToKey(code)

	u.email, _ = key.Lock([]byte(u.email), base.TwimoBase)
}

// Decode user's personal information with encr
func (u *user) Decode() {

	// KeyCode = last bit from comment code
	dividedComments := strings.SplitAfter(u.id, "-")
	code := dividedComments[len(dividedComments)-1]

	// Generate key
	key, _ := encr.StringToKey(code)

	email, _ := key.Unlock(u.email, base.TwimoBase)

	u.email = string(email)
}
