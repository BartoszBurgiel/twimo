package core

// Comment struct represents a comment
type Comment struct {
	Title   string
	Content string
	User    User

	// Id of the location
	Location string

	// Comment's ID
	ID string
}
