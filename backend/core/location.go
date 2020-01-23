package core

// Location struct represents a location
type Location struct {
	Name string

	// Store geological x and y coordinates
	Coords struct {
		X float64
		Y float64
	}

	Desc string

	HomePage string

	// Contains key to the comments in the database and comments objects
	Comments struct {
		Key      string
		Comments []Comment
	}

	Users string

	Ratings string

	ID string

	/*HELPER FIELDS*/

	// Link of the route on server
	LocationLink string
}
