package repository

import "twimo/backend/core"

// Repository is an interface a database needs to fulfill
// in order to be implemented into this API
type Repository interface {

	// Get a user with given ID
	GetUser(userID string) *core.User

	// Get a comment with given ID
	GetComment(commentID string) *core.Comment

	// Get a location with given ID
	GetLocation(locationID string) *core.Location

	// Get all comments of a location with given ID
	GetCommentsOfLocation(locationID string) []*core.Location
}
