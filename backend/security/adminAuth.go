package security

import (
	"io/ioutil"
	"twimo/backend/encr"
	"twimo/backend/encr/base"
)

// authAdminPassword checkes if admin password is valid
func authAdminPassword(adminPassword string) bool {
	// hash the password
	hashed := hashAdminPassword(adminPassword)
	// Read saved hashed original password
	originalPassword, _ := ioutil.ReadFile("../security/admin.pword")

	// if both passwords are equal
	if string(originalPassword) == hashed {
		return true
	}

	return false
}

// Hash admin's password
func hashAdminPassword(pword string) string {
	hash, _ := encr.Hash(pword, "40218790c028460c847ab633507350ff", base.Base85)
	return hash
}
