package core

// Comment struct represents a comment
type Comment struct {
	Title   string
	Content string
	User    User

	// Id of the location
	Location string

	// Rating of the location
	Rating int

	// Comment's ID
	ID string
}
