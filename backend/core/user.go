package core

// User struct represents a user
type User struct {
	Name     string
	Password string
	Email    string

	Comments struct {
		// Key to rows in Comment struct
		Key      string
		Comments []*Comment
	}

	// Key to rows in Location struct
	FavLocation string

	Ratings struct {
		// Key to rows in Ratings struct
		Key     string
		Ratings []*Rating
	}
	ID string
}
