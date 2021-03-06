package dummygenerator

import (
	"bytes"
)

// Generate random dummy data for twimo's database architecture
// return a buffer containing the querries
func Generate() bytes.Buffer {

	// Here all querries will be stored
	allQuerries := bytes.Buffer{}

	// Generate structs and querries for locations
	locationQuerries, locations := generateLocations()

	// Generate structs and querries for users
	userQuerries, users := generateUsers(locations[0])

	// Write generated querries for adding users
	allQuerries.WriteString(userQuerries)

	// Write generaetd querries for adding locations
	allQuerries.WriteString(locationQuerries)

	// Generate comment querries
	commentQuerries := generateComments(users, locations)

	// Write generated querries for adding comments
	allQuerries.WriteString(commentQuerries)

	// Generate features querries
	featureQuerries := generateFeatures(locations)

	// Write generated features querries
	allQuerries.WriteString(featureQuerries)

	return allQuerries
}

// GenerateEncoded random dummy data for twimo's database architecture
// return a buffer containing the querries
func GenerateEncoded() bytes.Buffer {
	// Here all querries will be stored
	allQuerries := bytes.Buffer{}
	// Generate structs and querries for locations
	locationQuerries, locations := generateLocations()

	// Generate structs and querries for users
	userQuerries, users := generateEncodedUsers(locations[0])

	// Write generated querries for adding users
	allQuerries.WriteString(userQuerries)

	// Write generaetd querries for adding locations
	allQuerries.WriteString(locationQuerries)

	// Decode users so that comment's title isn't hashed
	for i := 0; i < len(users); i++ {
		users[i].Decode()
	}

	// Generate comment querries
	commentQuerries := generateComments(users, locations)

	// Write generated querries for adding comments
	allQuerries.WriteString(commentQuerries)

	// Generate features querries
	featureQuerries := generateFeatures(locations)

	// Write generated features querries
	allQuerries.WriteString(featureQuerries)

	return allQuerries
}
