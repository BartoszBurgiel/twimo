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

	Webpage string

	Comments []Comment

	// Number of the users whose location is favorite
	Users int

	Rating float64

	// Price level
	Price int

	// Each attribute determins if a feature exists or not
	Features Features

	ID string
}
