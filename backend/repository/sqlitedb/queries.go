package sqlitedb

import (
	"fmt"
	"twimo/backend/core"
)

// GetUser returns a pointer to a user struct created from Db
func (r Repo) GetUser(userID string) *core.User {
	// Get  rows
	rows, err := r.db.Query(`SELECT name, password, email, comments, favlocation, ratings, id FROM users WHERE id = ? ; `, userID)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Iterate over rows

	// Variables
	var name, password, email, comments, favlocation, ratings, id string

	for rows.Next() {
		rows.Scan(&name, &password, &email, &comments, &favlocation, &ratings, &id)
	}

	// Build struct
	user := &core.User{
		Name:     name,
		Password: password,
		Email:    email,
		Comments: struct {
			Key      string
			Comments []*core.Comment
		}{
			Key:      comments,
			Comments: nil,
		},
		FavLocation: favlocation,
		Ratings: struct {
			Key     string
			Ratings []*core.Rating
		}{
			Key:     ratings,
			Ratings: nil,
		},
		ID: id,
	}

	fmt.Println("Chosen user: ", user)

	return user
}

// GetComment returns a pointer to a commment struct creater from Db
func (r Repo) GetComment(commentID string) *core.Comment {
	return nil
}

// GetLocation returns a pointer to a location struct creater from Db
func (r Repo) GetLocation(locationID string) *core.Location {
	return nil
}

// GetCommentsOfLocation returns a slice of pointers to comment structs from one location
func (r Repo) GetCommentsOfLocation(locationID string) []*core.Comment {
	return nil
}

// GetLocationsFavUsers returns a slice of pointers to user structs whose given location is the favorite
func (r Repo) GetLocationsFavUsers(locationID string) []*core.User {

	return nil
}

// GetLocationRatings returns a slice of pointers to rating structs
func (r Repo) GetLocationRatings(locationID string) []*core.Rating {

	return nil
}

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

// DeleteComment with a given ID from the database
func (r Repo) DeleteComment(commentID string) error {
	return nil
}

// DeleteRating with a given ID from the database
func (r Repo) DeleteRating(ratingID string) error {
	return nil
}
