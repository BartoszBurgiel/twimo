package dummygenerator

// Stores all data of the user needed to compose the querries
type user struct {
	name        string
	password    string
	email       string
	comments    string
	favlocation string
	ratings     string
	id          string
}

// Stores all data of the location needed to compose the querries
type location struct {
	name     string
	coordX   int
	coordY   int
	desc     string
	comments string
	users    string
	ratings  string
	id       string
}
