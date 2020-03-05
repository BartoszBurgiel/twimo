package repository

import (
	"fmt"
	"twimo/backend/core"
)

/*
	ALL INSERT INTO QUERRIES
*/

// AddComment to the database (all fields of struct must not be nil)
func (r Repo) AddComment(comment core.Comment) error {
	return nil
}

// SetAsFavoriteLocation the given location to a favorite location
func (r Repo) SetAsFavoriteLocation(locationID string, userID string) error {
	return nil
}

// AddUser to the database with the values from the struct passed as parameter
func (r Repo) AddUser(user core.User) error {
	_, err := r.db.Exec("INSERT INTO users VALUES(?, ?, ?, ?, ?) ;", user.Name, user.Password, user.Email, user.FavLocation, user.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("user has been added")
	return nil
}

// AddFeature to the database
func (r Repo) AddFeature(feature core.Location) error {
	return nil
}

// AddLocation to the database
func (r Repo) AddLocation(location core.Location) error {

	_, err := r.db.Exec("INSERT INTO locations VALUES(?, ?, ?, ?, ?, ?, ?) ;", location.Name, location.Coords.X, location.Coords.Y, location.Desc, location.Webpage, location.Price, location.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("location has been added")
	return nil

}
