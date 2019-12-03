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
	GetCommentsOfLocation(locationID string) []*core.Comment

	// Get all users that have given location as favorite
	GetLocationsFavUsers(locationID string) []*core.User

	// Get all ratings from the Db of given locations
	GetLocationRatings(locationID string) []*core.Rating

	// Add a comment by a user about a given location
	AddComment(locationID, userID, title, content string)

	// Add a rating about a location by a user
	AddRating(locationID, userID string, value int)

	// Set a given location as the favorite location of a user
	SetAsFavoriteLocation(locationID, userID string)

	// Delete a comment
	DeleteComment(commentID, userID string)

	// Delete a rating
	DeleteRating(ratingID, userID string)
}
