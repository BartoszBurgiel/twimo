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

// New location route with formatted link
func (l *LocationRoute) New(name string) {
	l.LocationName = name

	// Handle all umlauts
	name = strings.ReplaceAll(name, "ä", "+")
	name = strings.ReplaceAll(name, "ü", "*")
	name = strings.ReplaceAll(name, "ö", "$")
	name = "/" + strings.ReplaceAll(name, " ", "_")

	l.LocationLink = name

}
