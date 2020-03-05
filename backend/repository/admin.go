package repository

import (
	"fmt"
	"twimo/backend/core"
)

/*
ADMIN AND MAINTENENCE QUERRIES
*/

// GetAllUsersIDs from the database
func (r Repo) GetAllUsersIDs() (IDs []string, err error) {
	// Run querry
	rows, err := r.db.Query(`SELECT id FROM users ;`)
	if err != nil {
		fmt.Println(err)
		return IDs, err
	}

	// Schedule closing of the rows
	defer rows.Close()

	var tempID string

	// Iterate over rows
	for rows.Next() {
		err = rows.Scan(&tempID)
		if err != nil {
			return IDs, err
		}
		IDs = append(IDs, tempID)
	}
	return IDs, err
}

// GetAllLocations from the database
func (r Repo) GetAllLocations() (locations []core.Location, err error) {

	// Run querry
	rows, err := r.db.Query(`SELECT name, coordX, coordY, descr, webpage, id FROM locations ;`)
	if err != nil {
		fmt.Println(err)
		return locations, err
	}

	// Schedule closing of the rows
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {

		// Temporary variable for the location struct
		location := core.Location{}

		// Construct tempLocation
		err := rows.Scan(&location.Name, &location.Coords.X, &location.Coords.Y, &location.Desc, &location.Webpage, &location.ID)
		if err != nil {
			fmt.Println(err)
			return locations, err
		}

		// Fill remaining attributes

		// Get current location's comments
		comments, err := r.GetCommentsOfLocation(location.ID)
		if err != nil {
			fmt.Println(err)
			return locations, err
		}

		// Append to location struct
		location.Comments = comments

		// Get locations users
		users, err := r.GetLocationsFavUsersCount(location.ID)
		if err != nil {
			fmt.Println(err)
			return locations, err
		}

		// Append to location struct
		location.Users = users

		// Get location's average rating
		rating := r.GetLocationAvrRating(location.ID)
		location.Rating = rating

		// Get location's features
		features, err := r.GetLocationFeatures(location.ID)
		if err != nil {
			fmt.Println(err)
			return locations, err
		}

		// Append to location struct
		location.Features = features

		// Add tempLocation to locations
		locations = append(locations, location)
	}

	return locations, err
}

// GetAllUsers from the database
func (r Repo) GetAllUsers() (users []core.User, err error) {
	// Run querry
	rows, err := r.db.Query(`SELECT name, password, email, favlocation, id FROM users ;`)
	if err != nil {
		fmt.Println(err)
		return users, err
	}

	// Schedule closing of the rows
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {

		// Temporary variable for the user struct that will be build
		user := core.User{}

		// Construct tempuser
		err := rows.Scan(&user.Name, &user.Password, &user.Email, &user.FavLocation, &user.ID)
		if err != nil {
			fmt.Println(err)
			return users, err
		}

		// Decode user
		user.Decode()
		// Add tempuser to users
		users = append(users, user)
	}

	return users, err
}

// GetAllComments from the database
func (r Repo) GetAllComments() (comments []core.Comment, err error) {
	// Run querry
	rows, err := r.db.Query(`SELECT title, content, userID, locationID, id FROM comments ;`)
	if err != nil {
		fmt.Println(err)
		return comments, err
	}

	// Schedule closing of the rows
	defer rows.Close()

	var tempUserID, tempLocationID string

	// Iterate over rows
	for rows.Next() {

		// Temporary variable for the user struct that will be build
		comment := core.Comment{}

		// Construct tempuser
		err := rows.Scan(&comment.Title, &comment.Content, &tempUserID, &tempLocationID, &comment.ID)
		if err != nil {
			fmt.Println(err)
			return comments, err
		}

		// Fill remaining attributes

		// Get user
		user, err := r.GetUser(tempLocationID)
		if err != nil {
			fmt.Println(err)
			return comments, err
		}

		// Append to struct
		comment.User = user

		// Append to struct
		comment.Location = tempLocationID

		// Add tempuser to comments
		comments = append(comments, comment)
	}

	return comments, err
}

// GetLocationNames from the database and return these as a string
// slice
func (r Repo) GetLocationNames() (names []string, err error) {
	// Run querry
	rows, err := r.db.Query(`SELECT name FROM locations ;`)
	if err != nil {
		fmt.Println(err)
		return names, err
	}

	// Schedule closing of the rows
	defer rows.Close()

	var name string

	// Iterate over rows
	for rows.Next() {

		// Fill name
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
			return names, err
		}

		// Append name to names slice
		names = append(names, name)
	}

	return names, err
}

// GetLocationIDs returns the IDs of all locations in the database
func (r Repo) GetLocationIDs() (IDs []string, err error) {
	// Run querry
	rows, err := r.db.Query(`SELECT id FROM locations ;`)
	if err != nil {
		fmt.Println(err)
		return IDs, err
	}

	// Schedule closing of the rows
	defer rows.Close()

	var ID string

	// Iterate over rows
	for rows.Next() {

		// Fill name
		err := rows.Scan(&ID)
		if err != nil {
			fmt.Println(err)
			return IDs, err
		}

		// Append name to IDs slice
		IDs = append(IDs, "/"+ID)
	}

	return IDs, err
}

// UserExists in the database by couinting returning rows
// in the database
func (r Repo) UserExists(username string) (ok bool, err error) {
	// Run querry
	rows, err := r.db.Query(`SELECT id FROM users WHERE users.name = ? ;`, username)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// Schedule closing of the rows
	defer rows.Close()

	// counter of the rows
	cnt := 0

	// interate over rows
	for rows.Next() {
		cnt++
		if cnt > 0 {
			fmt.Println("User with the username", username, "was found in the database")
			return true, nil
		}
	}
	return false, nil
}
