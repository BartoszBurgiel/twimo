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

	Comments struct {
		// Key to rows in Comments
		Key      string
		Comments []*Comment
	}

	Ratings struct {
		// Key to rows in Ratings
		Key     string
		Ratings []*Rating
	}

	ID string
}
