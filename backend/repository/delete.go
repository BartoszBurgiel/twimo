package repository

import "fmt"

/*
ALL DELETE QUERRIES
*/

// DeleteComment with a given ID from the database
func (r Repo) DeleteComment(commentID string) error {
	// Get rating's ID
	rows, err := r.db.Query("SELECT ratingID from comments WHERE ID = ? ;", commentID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()
	ratingID := ""

	// iterate over rows
	for rows.Next() {
		err = rows.Scan(&ratingID)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// remove comment
	_, err = r.db.Exec("DELETE FROM comments WHERE ID = ? ; ", commentID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	r.DeleteRating(ratingID)
	return err
}

// DeleteRating with a given ID from the database
func (r Repo) DeleteRating(ratingID string) error {
	// remove rating
	_, err := r.db.Exec("DELETE FROM ratings WHERE ID = ? ; ", ratingID)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// DeleteUser from the datbase
func (r Repo) DeleteUser(userID string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE ID = ? ; ", userID)
	if err != nil {
		fmt.Println(err)
	}
	return err

}

// DeleteLocation from the database
func (r Repo) DeleteLocation(locationID string) error {

	// Delete from location table
	_, err := r.db.Exec("DELETE FROM locations WHERE ID = ? ; ", locationID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Delete all location's comments
	_, err = r.db.Exec("DELETE FROM comments WHERE locationID = ? ; ", locationID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Delete all location's features
	_, err = r.db.Exec("DELETE FROM features WHERE locationID = ? ; ", locationID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Delete all location's ratings
	_, err = r.db.Exec("DELETE FROM ratings WHERE locationID = ? ; ", locationID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Reset user's fav location
	_, err = r.db.Exec(`UPDATE users SET favlocation="" WHERE favlocation = ? ;`, locationID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// DeleteDish from the database
func (r Repo) DeleteDish(dishID string) error {
	return nil
}

// DeleteFeature from the database
func (r Repo) DeleteFeature(featureID string) error {
	return nil
}
