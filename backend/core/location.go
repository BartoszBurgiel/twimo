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

	Comments string

	Users string

	Ratings string

	ID string
}

// GenerateID based on the
func (l Location) GenerateID() string {
	return ""
}
