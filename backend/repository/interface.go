package repository

import "twimo/backend/core"

// Repository is an interface a database needs to fulfill
// in order to be implemented into this API
type Repository interface {

	// Get a user with given ID
	GetUser(userID string) (core.User, error)

	// Get comments given user has written
	GetUsersComments(userID string) ([]core.Comment, error)

	// Get user's name and ID
	GetUserDataForComment(userID string) (core.User, error)

	// Get user's data from the email
	GetUserFromEmail(userEmail string) (core.User, error)

	// Get user's data from the name
	GetUserFromName(userEmail string) (core.User, error)

	// Get a comment with given ID
	GetComment(commentID string) (core.Comment, error)

	// Get a location with given ID
	GetLocation(locationID string) (core.Location, error)

	// Get all data from location table that are needed on the location screen
	// criteria determine how they shall be sorted
	GetLocationsForList(criteria string) ([]core.Location, error)

	// Get a location with given name
	GetLocationFromName(locationName string) (core.Location, error)

	// Get all comments of a location with given ID
	GetCommentsOfLocation(locationID string) ([]core.Comment, error)

	// Get the average rating of a location
	GetLocationAvrRating(locationID string) (avr float64)

	// Get the number of users whose location is the favorite
	GetLocationsFavUsersCount(locationID string) (n int, err error)

	/*
		INSERT METHODS
	*/

	// Add a user to the database
	AddUser(core.User) error

	// Add a comment by a user about a given location
	AddComment(comment core.Comment) error

	// Add a feature to the database
	AddFeature(feature core.Location) error

	// Add one new location to the base
	AddLocation(location core.Location) error

	// Set a given location as the favorite location of a user
	SetAsFavoriteLocation(locationID, userID string) error

	/*
		DELETE METHODS
	*/

	// Delete a comment from the data base with given ID
	DeleteComment(commentID string) error

	// Delete a rating from the data base with given ID
	DeleteRating(ratingID string) error

	// Delete an user from the data base with given ID
	DeleteUser(userID string) error

	// Delete a location from the data base with given ID
	DeleteLocation(locationID string) error

	// Delete a feature from the data base with given ID
	DeleteFeature(featureID string) error

	/*
		ADMIN AND MAINTENECE QUERRIES
	*/

	// Get all users from the database
	GetAllUsers() ([]core.User, error)

	// Get all locations from the database
	GetAllLocations() ([]core.Location, error)

	// Get all comments from the database
	GetAllComments() ([]core.Comment, error)

	// Get all names of the locations
	GetLocationNames() ([]string, error)

	// Get all IDs of the locations
	GetLocationIDs() ([]string, error)

	// Get all IDs of the users
	GetAllUsersIDs() ([]string, error)
}
