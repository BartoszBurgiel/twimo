package repository

import "twimo/backend/core"

/*
	ALL INSERT INTO QUERRIES
*/

// AddComment to the database (all fields of struct must not be nil)
func (r Repo) AddComment(comment core.Comment) error {
	return nil
}

// AddRating to the database (all fields of struct must not be nil)
func (r Repo) AddRating(rating core.Rating) error {
	return nil
}

// SetAsFavoriteLocation the given location to a favorite location
func (r Repo) SetAsFavoriteLocation(locationID string, userID string) error {
	return nil
}

// AddUser to the database with the values from the struct passed as parameter
func (r Repo) AddUser(user core.User) error {
	return nil
}

// AddDish to the datbaase
func (r Repo) AddDish(dish core.Dish) error {
	return nil
}

// AddFeature to the database
func (r Repo) AddFeature(feature core.Feature) error {
	return nil
}

// AddLocation to the database
func (r Repo) AddLocation(location core.Location) error {
	return nil
}
