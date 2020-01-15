package repository

/*
	ALL SELECT QUERRIES
*/

import (
	"fmt"
	"twimo/backend/core"
)

// GetUser returns a pointer to a user struct created from Db
func (r Repo) GetUser(userID string) (user core.User, err error) {

	// Set userID
	user.ID = userID

	// Get  rows
	rows, err := r.db.Query(`SELECT name, password, email, comments, favlocation, ratings FROM users WHERE id = ? ; `, userID)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&user.Name, &user.Password, &user.Email, &user.Comments, &user.FavLocation, &user.Ratings)
	}

	// Check if name is empty -> no users found
	if user.Name == "" {
		return user, fmt.Errorf("No users found with id %s", userID)
	}

	// Decode user
	user.Decode()

	return user, err
}

// GetUsersComments from the database
func (r Repo) GetUsersComments(userID string) (comments []core.Comment, err error) {
	return comments, err
}

// GetUsersRatings from the database
func (r Repo) GetUsersRatings(userID string) (ratings []core.Rating, err error) {
	return ratings, err
}

// GetComment returns a pointer to a commment struct creater from Db
func (r Repo) GetComment(commentID string) (comment core.Comment, err error) {
	// Set comment;s id
	comment.ID = commentID

	// Get  rows
	rows, err := r.db.Query(`SELECT title, content, userID, locationID FROM comments WHERE id = ? ; `, commentID)
	if err != nil {
		fmt.Println(err)
		return comment, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&comment.Title, &comment.Content, &comment.User, &comment.Location)
	}

	// Check if title is empty -> comment not found
	if comment.Title == "" {
		return comment, fmt.Errorf("Comment with id %s not found in the database", commentID)
	}

	return comment, err
}

// GetLocation returns a pointer to a location struct creater from Db
func (r Repo) GetLocation(locationID string) (location core.Location, err error) {

	location.ID = locationID

	// Get  rows
	rows, err := r.db.Query(`SELECT name, coordX, coordY, descr, comments, users, ratings FROM locations WHERE id = ?  ; `, locationID)
	if err != nil {
		fmt.Println(err)
		return location, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&location.Name, &location.Coords.X, &location.Coords.Y, &location.Desc, &location.Comments, &location.Users, &location.Ratings)
	}

	// Check if title is empty -> comment not found
	if location.Name == "" {
		return location, fmt.Errorf("Location with id %s not found in the database", locationID)
	}

	// // Set location's id
	// location.ID = locationID

	return location, err
}

// GetLocationsAsLink returns all locations' data so that they be displayed as an <a> tag: name and link
func (r Repo) GetLocationsAsLink() (locations []core.LocationRoute, err error) {

	// Get  rows
	rows, err := r.db.Query(`SELECT name FROM locations ; `)
	if err != nil {
		fmt.Println(err)
		return locations, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		tempLocationRoute := core.LocationRoute{}
		name := ""

		rows.Scan(&name)

		// Check if name is empty -> comment not found
		if name == "" {
			return locations, fmt.Errorf("No locations found in the database")
		}

		// Create route and append to the locations slice
		tempLocationRoute.New(name)

		locations = append(locations, tempLocationRoute)
	}

	return locations, err
}

// GetLocationFromName return all rows from a the database of a location with a given name
func (r Repo) GetLocationFromName(locationName string) (location core.Location, err error) {

	// Get  rows
	rows, err := r.db.Query(`SELECT name, coordX, coordY, descr, comments, users, ratings FROM locations WHERE name = ?  ; `, locationName)
	if err != nil {
		fmt.Println(err)
		return location, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&location.Name, &location.Coords.X, &location.Coords.Y, &location.Desc, &location.Comments, &location.Users, &location.Ratings)
	}

	// Check if title is empty -> comment not found
	if location.Name == "" {
		return location, fmt.Errorf("Location with id %s not found in the database", locationName)
	}

	// // Set location's id
	// location.ID = locationID

	return location, err
}

// GetCommentsOfLocation returns a slice of pointers to comment structs from one location
func (r Repo) GetCommentsOfLocation(locationID string) (comments []core.Comment, err error) {
	return comments, err
}

// GetLocationsFavUsers returns a slice of pointers to user structs whose given location is the favorite
func (r Repo) GetLocationsFavUsers(locationID string) (users []core.User, err error) {

	return users, err
}

// GetLocationRatings returns a slice of pointers to rating structs
func (r Repo) GetLocationRatings(locationID string) (ratings []core.Rating, err error) {

	return ratings, err
}

// GetLocationAvrRating from the database
func (r Repo) GetLocationAvrRating(locationID string) (avr float64) {
	return avr
}

// GetLocationsDishes from the database
func (r Repo) GetLocationsDishes(locationID string) (dishes []core.Dish, err error) {
	return dishes, err
}

// GetLocationsFeatures from the database
func (r Repo) GetLocationsFeatures(locationID string) (features []core.Feature, err error) {
	return features, err
}
