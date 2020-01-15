package core

import "strings"

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

	Comments string

	Users string

	Ratings string

	ID string
}

// LocationRoute contains all information about a location so that it can be displayed
// as a <a> link
type LocationRoute struct {
	// Original name
	LocationName string
	// Link of the route on server
	LocationLink string
}

// New locationRoute created from the name
func (l *LocationRoute) New(locationName string) {
	l.LocationName = locationName
	l.LocationLink = "/" + strings.ReplaceAll(locationName, " ", "_")

	// Handle all umlauts
	l.LocationLink = strings.ReplaceAll(l.LocationLink, "ä", "ae")
	l.LocationLink = strings.ReplaceAll(l.LocationLink, "ü", "ue")
	l.LocationLink = strings.ReplaceAll(l.LocationLink, "ö", "oe")
}
