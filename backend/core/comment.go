package core

// Comment struct represents a comment
type Comment struct {
	Title   string
	Content string

	// User's ID and user object
	User struct {
		Key  string
		User User
	}

	// Location's ID
	Location string

	// Comment's ID
	ID string
}
