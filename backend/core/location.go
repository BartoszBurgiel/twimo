package core

// Location struct represents a location
type Location struct {
	Name string

	// Store geological x and y coordinates
	Coords struct {
		X float64
		Y float64
	}

	// Description
	Desc string

	// Key to rows in Comments
	Comments string

	// Key to rows in Users
	Users string

	// Key to rows in Ratings
	Ratings string
	ID      string
}
