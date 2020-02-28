package repository

import (
	"encoding/json"
	"io/ioutil"
	"twimo/backend/core"
)

// keep the most complex (more that 1 table involved
// or some kind of parameter)
// and most important queries
// and their parameter in one place
type vars struct {
	// Path to vars file
	path string

	LocationForListLimit    int `json:"LocationForListLimit"`
	CommentsOfLocationLimit int `json:"CommentsOfLocationLimit"`

	// pointers to all fields needed to display
	// a location on the location list
	locationListPointers []interface{}
}

// Initiate the vars
func (v *vars) init(pathToJSON string) {
	v.path = pathToJSON

	// Read JSON
	conf, err := ioutil.ReadFile(pathToJSON)
	if err != nil {
		panic(err)
	}

	// parse values from the json to struct
	json.Unmarshal(conf, v)
	if err != nil {
		panic(err)
	}
}

// get all needed pointers of the location in order to
// call the scan function
func (v *vars) getLocationPointer(loc *core.Location) {
	v.locationListPointers = []interface{}{&loc.Name, &loc.Desc, &loc.ID, &loc.Price, &loc.Rating}
}

// get a query for the locations according to the criteria
func getLocationListQuery(criteria string) string {
	switch criteria {
	case "rating":
		return `SELECT locations.name, locations.descr, locations.id, locations.price, AVG(ratings.value)
		FROM locations
		LEFT JOIN ratings
		ON locations.id = ratings.locationID
		ORDER BY AVG(ratings.value) DESC
		LIMIT %d`
	case "nFavUsers":
		return `SELECT locations.name, locations.descr, locations.id, locations.price, AVG(ratings.value)
		FROM locations
		LEFT JOIN users
		ON locations.id = users.favlocation
		ORDER BY COUNT(users.favlocation) DESC
		LIMIT %d
	  `
	case "price":
		return `SELECT locations.name, locations.descr, locations.id, locations.price, AVG(ratings.value)
		FROM locations 
		ORDER BY locations.price DESC
		LIMIT %d`
	}
	return ""
}
