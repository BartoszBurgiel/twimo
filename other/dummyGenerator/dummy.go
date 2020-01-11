package dummygenerator

import (
	"bytes"
)

// Generate random dummy data for twimo's database architecture
// return a buffer containing the querries
func Generate() bytes.Buffer {

	// Here all querries will be stored
	allQuerries := bytes.Buffer{}

	// Generate structs and querries for users
	userQuerries, users := generateUsers()

	// Write generated querries for adding users
	allQuerries.WriteString(userQuerries)

	// Generate structs and querries for locations
	locationQuerries, locations := generateLocations()

	// Write generaetd querries for adding locations
	allQuerries.WriteString(locationQuerries)

	// Generate rating queries
	ratingQueries := generateRatings(users, locations)

	// Write generated querries for adding ratings
	allQuerries.WriteString(ratingQueries)

	// Generate comment querries
	commentQuerries := generateComments(users, locations)

	// Write generated querries for adding comments
	allQuerries.WriteString(commentQuerries)

	// Generate fav location querries
	favlocationQuerries := generateFavLocations(users, locations)

	// Write generated querries for adding fav locations
	allQuerries.WriteString(favlocationQuerries)

	// Generate features querries
	featureQuerries := generateFeatures(locations)

	// Write generated features querries
	allQuerries.WriteString(featureQuerries)

	// Generate dishes querries
	dishesQuerries := generateDishes(locations)

	// Write generated dishes querries
	allQuerries.WriteString(dishesQuerries)

	return allQuerries
}

// GenerateEncoded random dummy data for twimo's database architecture
// return a buffer containing the querries
func GenerateEncoded() bytes.Buffer {
	// Here all querries will be stored
	allQuerries := bytes.Buffer{}

	// Generate structs and querries for users
	userQuerries, users := generateEncodedUsers()

	// Write generated querries for adding users
	allQuerries.WriteString(userQuerries)

	// Generate structs and querries for locations
	locationQuerries, locations := generateLocations()

	// Write generaetd querries for adding locations
	allQuerries.WriteString(locationQuerries)

	// Generate rating queries
	ratingQueries := generateRatings(users, locations)

	// Write generated querries for adding ratings
	allQuerries.WriteString(ratingQueries)

	// Decode users so that comment's title isn't hashed
	for i := 0; i < len(users); i++ {
		users[i].Decode()
	}

	// Generate comment querries
	commentQuerries := generateComments(users, locations)

	// Write generated querries for adding comments
	allQuerries.WriteString(commentQuerries)

	// Generate fav location querries
	favlocationQuerries := generateFavLocations(users, locations)

	// Write generated querries for adding fav locations
	allQuerries.WriteString(favlocationQuerries)

	// Generate features querries
	featureQuerries := generateFeatures(locations)

	// Write generated features querries
	allQuerries.WriteString(featureQuerries)

	// Generate dishes querries
	dishesQuerries := generateDishes(locations)

	// Write generated dishes querries
	allQuerries.WriteString(dishesQuerries)

	return allQuerries
}
