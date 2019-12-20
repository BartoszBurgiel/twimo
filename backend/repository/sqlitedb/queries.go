package sqlitedb

import (
	"fmt"
	"twimo/backend/core"
)

// GetUser returns a pointer to a user struct created from Db
func (r Repo) GetUser(userID string) (*core.User, error) {
	// Get  rows
	rows, err := r.db.Query(`SELECT name, password, email, comments, favlocation, ratings, id FROM users WHERE id = ? ; `, userID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows

	// Variables
	var name, password, email, comments, favlocation, ratings, id string

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&name, &password, &email, &comments, &favlocation, &ratings, &id)
	}

	// Check if name is empty -> no users found
	if name == "" {
		return nil, fmt.Errorf("No users found with id %s", userID)
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

	return user, nil
}

// GetComment returns a pointer to a commment struct creater from Db
func (r Repo) GetComment(commentID string) (*core.Comment, error) {
	// Get  rows
	rows, err := r.db.Query(`SELECT title, content, userID, locationID, id FROM comments WHERE id = ? ; `, commentID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Schedule closing of the db
	defer rows.Close()

	// Define variables
	var title, content, userID, locationID, id string

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&title, &content, &userID, &locationID, &id)
	}

	// Check if title is empty -> comment not found
	if title == "" {
		return nil, fmt.Errorf("Comment with id %s not found in the database", commentID)
	}

	// Construct comment struct
	comment := core.Comment{
		Title:    title,
		Content:  content,
		Location: locationID,
		User:     userID,
		ID:       commentID,
	}

	return &comment, nil
}

// GetLocation returns a pointer to a location struct creater from Db
func (r Repo) GetLocation(locationID string) (*core.Location, error) {
	// Get  rows
	rows, err := r.db.Query(`SELECT name, coordX, coordY, desc, comments, users, ratings FROM locations WHERE id = ?  ; `, locationID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Schedule closing of the db
	defer rows.Close()

	// Define variables
	var name, desc, comments, users, ratings string
	var coordX, coordY float64

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&name, &coordX, &coordY, &desc, &comments, &users, &ratings)
	}

	// Check if title is empty -> comment not found
	if name == "" {
		return nil, fmt.Errorf("Location with id %s not found in the database", locationID)
	}

	// Construct location struct
	location := core.Location{
		Name: name,
		Desc: desc,
		Comments: struct {
			Key      string
			Comments []*core.Comment
		}{
			Key:      comments,
			Comments: nil,
		},
		Coords: struct {
			X float64
			Y float64
		}{
			X: coordX,
			Y: coordY,
		},
		Ratings: struct {
			Key     string
			Ratings []*core.Rating
		}{
			Key:     ratings,
			Ratings: nil,
		},
		ID: locationID,
	}

	return &location, nil
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
