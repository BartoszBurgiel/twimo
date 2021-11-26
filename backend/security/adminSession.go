package security

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte(uuid.New().String())
	store = sessions.NewCookieStore(key)
)

// AuthAdminSession authenticates admin session
// and checks if the admin has had logged in
func AuthAdminSession(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "admin-login-session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}

	return true
}

// AdminLogin hosts admin login page and checks the password
func AdminLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "admin-login-session")

	// Fetch password
	pword := r.FormValue("pword")
	fmt.Println(pword)

	if pword != "" {

		// Set user as authenticated
		session.Values["authenticated"] = authAdminPassword(pword)
	} else {

		// if no password entered -> kihmo unvalid
		// and just set the authenticated value to false
		session.Values["authenticated"] = false
	}
	session.Save(r, w)
}

// CheckKillSession checks if the admin destroyed his session
// and renders the authentication to unvalid
func CheckKillSession(w http.ResponseWriter, r *http.Request) {
	killSession := r.FormValue("killSession")

	// If killSession has a value ->
	// if admin insisted on destroying the session
	if killSession != "" {
		fmt.Println("yo!")
		session, _ := store.Get(r, "admin-login-session")

		// Revoke users authentication
		session.Values["authenticated"] = false
		session.Save(r, w)
	}
}
