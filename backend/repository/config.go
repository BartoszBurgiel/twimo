package repository

import (
	"encoding/json"
	"io/ioutil"
)

// keep the most complex (more that 1 table involved
// or some kind of parameter)
// and most important queries
// and their parameter in one place
type config struct {
	// Path to config file
	path string

	LocationForListLimit    int `json:"LocationForListLimit"`
	CommentsOfLocationLimit int `json:"CommentsOfLocationLimit"`
}

// Initiate the config
func (c *config) init(pathToJSON string) {
	c.path = pathToJSON

	// Read JSON
	conf, err := ioutil.ReadFile(pathToJSON)
	if err != nil {
		panic(err)
	}

	// parse values from the json to struct
	json.Unmarshal(conf, c)
	if err != nil {
		panic(err)
	}
}
