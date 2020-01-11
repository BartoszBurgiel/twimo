package repository

import "twimo/backend/core"

// Repository is an interface a database needs to fulfill
// in order to be implemented into this API
type Repository interface {

	// Get a user with given ID
	GetUser(userID string) (*core.User, error)

	// Get comments given user has written
	GetUsersComments(userID string) ([]*core.Comment, error)

	// Get ratings user has given
	GetUsersRatings(userID string) ([]*core.Rating, error)

	// Get a comment with given ID
	GetComment(commentID string) (*core.Comment, error)

	// Get a location with given ID
	GetLocation(locationID string) (*core.Location, error)

	// Get all comments of a location with given ID
	GetCommentsOfLocation(locationID string) ([]*core.Comment, error)

	// Get all users that have given location as favorite
	GetLocationsFavUsers(locationID string) ([]*core.User, error)

	// Get all ratings from the Db of given locations
	GetLocationRatings(locationID string) ([]*core.Rating, error)

	// Get the average of all ratings of the location
	GetLocationAvrRating(locationID string) float64

	// Get all of the dishes that a location offers
	GetLocationsDishes(locationID string) ([]*core.Dish, error)

	// Get all of the features that a location offers
	GetLocationsFeatures(locationID string) ([]*core.Feature, error)

	// Add a user to the database
	AddUser(core.User) error

	// Add a comment by a user about a given location
	AddComment(comment core.Comment) error

	// Add a rating about a location by a user
	AddRating(rating core.Rating) error

	// Add a dish to the database
	AddDish(dish core.Dish) error

	// Add a feature to the database
	AddFeature(feature core.Feature) error

	// Add one new location to the base
	AddLocation(location core.Location) error

	// Set a given location as the favorite location of a user
	SetAsFavoriteLocation(locationID, userID string) error

	// Delete a comment from the data base with given ID
	DeleteComment(commentID string) error

	// Delete a rating from the data base with given ID
	DeleteRating(ratingID string) error

	// Delete an user from the data base with given ID
	DeleteUser(userID string) error

	// Delete a location from the data base with given ID
	DeleteLocation(locationID string) error

	// Delete a dish from the data base with given ID
	DeleteDish(dishID string) error

	// Delete a feature from the data base with given ID
	DeleteFeature(featureID string) error

	/*
		ADMIN AND MAINTENECE QUERRIES
	*/

	// Get all users from the database
	GetAllUsers() ([]*core.User, error)

	// Get all locations from the database
	GetAllLocations() ([]*core.Location, error)

	// Get all comments from the database
	GetAllComments() ([]*core.Comment, error)
}
