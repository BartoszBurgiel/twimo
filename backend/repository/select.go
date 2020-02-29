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
	rows, err := r.db.Query(`SELECT name, password, email, favlocation FROM users WHERE id = ? ; `, userID)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&user.Name, &user.Password, &user.Email, &user.FavLocation)
	}

	// Check if name is empty -> no users found
	if user.Name == "" {
		return user, fmt.Errorf("No users found with id %s", userID)
	}

	// Decode user
	user.Decode()

	return user, err
}

// GetUserDataForComment that will be displayed in frontend
// - full name + id (for avatar)
func (r Repo) GetUserDataForComment(userID string) (user core.User, err error) {

	// Get user from the db
	user, err = r.GetUser(userID)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	// Set all unnecessary fields to ""
	// just name and ID can't be pulled from the database, because several more fiels are needed
	// to decode user's data
	user.Email = ""
	user.Password = ""
	user.FavLocation = ""

	return user, err
}

// GetUserFromEmail returns a user struct with all user's data
func (r Repo) GetUserFromEmail(userEmail string) (user core.User, err error) {

	// Set userID
	user.Email = userEmail

	// Get  rows
	rows, err := r.db.Query(`SELECT name, password, favlocation, id FROM users WHERE email = ? ; `, userEmail)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&user.Name, &user.Password, &user.FavLocation, &user.ID)
	}

	// Check if name is empty -> no users found
	if user.Name == "" {
		return user, fmt.Errorf("No users found with email %s", userEmail)
	}

	// Decode user
	user.Decode()

	return user, err
}

// GetUserFromName returns a user struct with all user's data
func (r Repo) GetUserFromName(userName string) (user core.User, err error) {

	// Set userID
	user.Name = userName

	// Get  rows
	rows, err := r.db.Query(`SELECT email, password, favlocation, id FROM users WHERE name = ? ; `, userName)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&user.Email, &user.Password, &user.FavLocation, &user.ID)
	}

	// Check if name is empty -> no users found
	if user.Email == "" {
		return user, fmt.Errorf("No users found with name %s", userName)
	}

	// Decode user
	user.Decode()

	return user, err
}

// GetUsersComments from the database
func (r Repo) GetUsersComments(userID string) (comments []core.Comment, err error) {
	return comments, err
}

// GetComment returns a pointer to a commment struct creater from Db
func (r Repo) GetComment(commentID string) (comment core.Comment, err error) {
	// Set comment;s id
	comment.ID = commentID

	// Get  rows
	rows, err := r.db.Query(`SELECT title, content, userID, locationID, ratingID FROM comments WHERE id = ? ; `, commentID)
	if err != nil {
		fmt.Println(err)
		return comment, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	userID := ""
	for rows.Next() {
		rows.Scan(&comment.Title, &comment.Content, &userID, &comment.Location, &comment.Rating)
	}

	// Get user struct
	user, err := r.GetUser(userID)
	if err != nil {
		fmt.Println(err)
		return comment, err
	}

	comment.User = user

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
	rows, err := r.db.Query(`SELECT name, coordX, coordY, descr, webpage FROM locations WHERE id = ?  ; `, locationID)
	if err != nil {
		fmt.Println(err)
		return location, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&location.Name, &location.Coords.X, &location.Coords.Y, &location.Desc, &location.Webpage)
	}

	// Check if title is empty -> comment not found
	if location.Name == "" {
		return location, fmt.Errorf("Location with id %s not found in the database", locationID)
	}

	// Set reminding attributes of the location struct
	// Comments
	comments, err := r.GetCommentsOfLocation(locationID)
	if err != nil {
		fmt.Println(err)
		return location, err
	}
	location.Comments = comments

	// Users
	users, err := r.GetLocationsFavUsersCount(locationID)
	if err != nil {
		fmt.Println(err)
		return location, err
	}
	location.Users = users

	// Rating
	rating := r.GetLocationAvrRating(locationID)
	location.Rating = rating

	// Pricing
	pricing := r.GetLocationsPrice(locationID)
	location.Price = pricing

	// Features
	features, err := r.GetLocationFeatures(locationID)
	if err != nil {
		fmt.Println(err)
		return location, err
	}
	location.Features = features

	return location, err
}

// GetLocationFeatures from the database
func (r Repo) GetLocationFeatures(locationID string) (features core.Features, err error) {
	// Get  rows
	rows, err := r.db.Query(`SELECT coffee, wifi, power, music, food FROM features WHERE locationID = ?  ; `, locationID)
	if err != nil {
		fmt.Println(err)
		return features, err
	}

	// Schedule closing of the db
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&features.Coffee, &features.Wifi, &features.Power, &features.Music, &features.Food)
	}

	return features, err
}

// GetLocationFromName return all rows from a the database of a location with a given name
func (r Repo) GetLocationFromName(locationName string) (location core.Location, err error) {

	// Get  rows
	rows, err := r.db.Query(`SELECT name, coordX, coordY, descr, webpage, id FROM locations WHERE name = ?  ; `, locationName)
	if err != nil {
		fmt.Println(err)
		return location, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&location.Name, &location.Coords.X, &location.Coords.Y, &location.Desc, &location.Webpage, &location.ID)
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
	// Prepare query
	query := fmt.Sprintf(`SELECT title, content, userID, locationID, ratingID, id FROM comments WHERE locationID = ? LIMIT %d ; `, r.vars.CommentsOfLocationLimit)

	// Get  rows
	rows, err := r.db.Query(query, locationID)
	if err != nil {
		fmt.Println(err)
		return comments, err
	}

	// Schedule closing of the db
	defer rows.Close()

	// Temporary comment that will be filled by scan
	var comment core.Comment
	var tempUserID, tempLocationID, tempRatingID string

	// Iterate over rows
	for rows.Next() {
		rows.Scan(&comment.Title, &comment.Content, &tempUserID, &tempLocationID, &tempRatingID, &comment.ID)

		// Check if title is empty -> comment not found
		if comment.Title == "" {
			return comments, fmt.Errorf("No comments found ")
		}

		/*
			Append remaining data to comment structs
		*/

		// Pull user from the database
		user, err := r.GetUserDataForComment(tempUserID)
		if err != nil {
			fmt.Println(err)
			return comments, err
		}

		// Append user to comment struct
		comment.User = user

		// Append location to comment struct
		comment.Location = tempLocationID

		// Set rating
		rating, err := r.GetRating(tempRatingID)
		if err != nil {
			fmt.Println(err)
			return comments, err
		}

		comment.Rating = rating

		// Append to comments
		comments = append(comments, comment)
	}

	return comments, err
}

// GetLocationsFavUsersCount returns the number of location's favorite users
func (r Repo) GetLocationsFavUsersCount(locationID string) (n int, err error) {

	// Get rows from the database
	rows, err := r.db.Query(`SELECT COUNT(name) FROM users WHERE favLocation = ? ;`, locationID)
	if err != nil {
		fmt.Println(err)
		return n, err
	}

	// Iterate over rows
	for rows.Next() {
		err = rows.Scan(&n)
		if err != nil {
			fmt.Println(err)
			return n, err
		}
	}

	return n, err
}

// GetLocationAvrRating from the database
func (r Repo) GetLocationAvrRating(locationID string) (avr float64) {

	// Get rows from the database
	rows, err := r.db.Query(`SELECT AVG(value) FROM ratings WHERE locationID = ? ;`, locationID)
	if err != nil {
		fmt.Println(err)
		return avr
	}

	// Iterate over rows
	for rows.Next() {
		err = rows.Scan(&avr)
		if err != nil {
			fmt.Println(err)
			return avr
		}
	}

	return avr
}

// GetLocationsPrice returns the price of the location with the given ID
func (r Repo) GetLocationsPrice(locationID string) (n int) {

	// Get rows from the database
	rows, err := r.db.Query(`SELECT price FROM locations WHERE id = ? ;`, locationID)
	if err != nil {
		fmt.Println(err)
		return n
	}

	// Iterate over rows
	for rows.Next() {
		err = rows.Scan(&n)
		if err != nil {
			fmt.Println(err)
			return n
		}
	}

	return n
}

// GetRating with the given ID from the database
func (r Repo) GetRating(ratingID string) (rating int, err error) {

	// Get rows from the database
	rows, err := r.db.Query(`SELECT value FROM ratings WHERE id = ? ;`, ratingID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Iterate over rows
	for rows.Next() {

		// Fill the rating variable
		err = rows.Scan(&rating)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	}
	return rating, nil
}

// GetLocationsForList returns all data that is displayed on the
// list screen
func (r Repo) GetLocationsForList(criteria string) (locations []core.Location, err error) {
	// determine query based on the criteria
	query := getLocationListQuery(criteria)

	// Prepare query
	query = fmt.Sprintf(query, r.vars.LocationForListLimit)

	// Get rows from the database
	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return locations, err
	}

	// Iterate over rows
	for rows.Next() {
		tempLoc := core.Location{}

		r.vars.getLocationPointer(&tempLoc)
		err = rows.Scan(r.vars.locationListPointers...)
		if err != nil {
			fmt.Println(err)
			return locations, err
		}

		// Append to locations
		locations = append(locations, tempLoc)
	}

	return locations, err
}
