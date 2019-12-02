package core

// User struct represents a user
type User struct {
	Name     string
	Password string
	Email    string

	// Key to rows in Comment struct
	Comments string

	// Key to rows in Location struct
	FavLocation string

	// Key to rows in Ratings struct
	Ratings string
	ID      string
}
