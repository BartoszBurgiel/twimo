package repository

import "twimo/backend/core"

// Repository is an interface a database needs to fulfill
// in order to be implemented into this API
type Repository interface {

	// Get a user with given ID
	GetUser(userID string) (*core.User, error)

	// Get a comment with given ID
	GetComment(commentID string) (*core.Comment, error)

	// Get a location with given ID
	GetLocation(locationID string) (*core.Location, error)

	// Get all comments of a location with given ID
	GetCommentsOfLocation(locationID string) []*core.Comment

	// Get all users that have given location as favorite
	GetLocationsFavUsers(locationID string) []*core.User

	// Get all ratings from the Db of given locations
	GetLocationRatings(locationID string) []*core.Rating

	// Add a comment by a user about a given location
	AddComment(comment core.Comment) error

	// Add a rating about a location by a user
	AddRating(rating core.Rating) error

	// Set a given location as the favorite location of a user
	SetAsFavoriteLocation(locationID, userID string) error

	// Delete a comment
	DeleteComment(commentID string) error

	// Delete a rating
	DeleteRating(ratingID string) error
}
