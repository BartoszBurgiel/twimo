package dummygenerator

import (
	"fmt"
	"math/rand"
	"time"
	"twimo/backend/encr"
	"twimo/backend/encr/base"

	"github.com/google/uuid"
)

/*
	HERE ARE ALL GENERATE FUNCTIONS FOR EVERY TABLE

	EACH FUNCTION RETURNS BOTH THE QUERRIES AS STRING
	AND THE STRUCT SO THAT THE DATA USED TO WRITE THE
	QUERRY CAN BE REUSED
*/

// Generate user querries
func generateUsers(location location) (string, []user) {

	// Array storing all names
	names := []string{
		"Louis",
		"Bartosz",
		"Jannik",
		"Kai",
		"Tim",
		"Bettina",
		"Aleksandra",
		"Adalbert",
	}

	// Array stroing all surnames
	surnames := []string{
		"Beul",
		"Schneider",
		"Müller",
		"Spohr",
		"Stenz",
		"Sauer",
	}

	// Array storing all user structs
	users := []user{}

	/*
	 Generate users
	 Iterate over both arrays and generate all possible users
	*/

	// Iterate over names
	for i := 0; i < len(names); i++ {

		// Iterate over surnames
		for j := 0; j < len(surnames); j++ {

			// Assemble field variables
			fullName := names[i] + " " + surnames[j]

			// first letter of name+.+surname+@mail.com
			email := names[i] + "." + surnames[j] + "@main.com"
			password := names[i] + surnames[j] + "69"

			/* BARTOSZ'S PART */
			// hash password
			_, salt, _ := encr.Salt(fullName, base.TwimoBase)
			password, err := encr.Hash(password, salt, base.TwimoBase)
			if err != nil {
				panic(err)
			}

			id := uuid.New().String()

			// Build temporary user struct
			tempUser := user{
				name:        fullName,
				email:       email,
				password:    password,
				favlocation: location.id,
				id:          id,
			}

			// Add temp user to the array of users
			users = append(users, tempUser)
		}
	}

	/*
	 Write querries
	*/

	querries := ""

	// Iterate over users
	for _, user := range users {

		// Setup querry with 'printf' - like casting
		querries += fmt.Sprintf(addNewUser, user.name, user.password, user.email, user.favlocation, user.id)

		// Add line break
		querries += "\n"
	}

	return querries, users
}

// Generate user querries
func generateEncodedUsers(location location) (string, []user) {

	// Array storing all names
	names := []string{
		"Louis",
		"Bartosz",
		"Jannik",
		"Kai",
		"Lukas",
		"Luca",
		"Adalbert",
	}

	// Array stroing all surnames
	surnames := []string{
		"Beul",
		"Langscheid",
		"Stenz",
		"Sauer",
	}

	// Array storing all user structs
	users := []user{}

	/*
	 Generate users
	 Iterate over both arrays and generate all possible users
	*/

	// Iterate over names
	for i := 0; i < len(names); i++ {

		// Iterate over surnames
		for j := 0; j < len(surnames); j++ {

			// Assemble field variables
			fullName := names[i] + " " + surnames[j]

			// first letter of name+.+surname+@mail.com
			email := names[i] + "." + surnames[j] + "@main.com"
			password := names[i] + surnames[j] + "69"

			/* BARTOSZ'S PART */
			// hash password
			_, salt, _ := encr.Salt(fullName, base.TwimoBase)
			password, err := encr.Hash(password, salt, base.TwimoBase)
			if err != nil {
				panic(err)
			}

			id := uuid.New().String()

			// Build temporary user struct
			tempUser := user{
				name:        fullName,
				email:       email,
				password:    password,
				favlocation: location.id,
				id:          id,
			}

			// Encode tempUser
			tempUser.Encode()

			// Add temp user to the array of users
			users = append(users, tempUser)
		}
	}

	/*
	 Write querries
	*/

	querries := ""

	// Iterate over users
	for _, user := range users {

		// Setup querry with 'printf' - like casting
		querries += fmt.Sprintf(addNewUser, user.name, user.password, user.email, user.favlocation, user.id)

		// Add line break
		querries += "\n"
	}

	return querries, users
}

// generate location querries
func generateLocations() (string, []location) {

	// Array storing names for the location
	names := []string{
		"Betties",
		"Westerburger",
		"Lukas",
	}

	// Array storing place types (?) of the locaton
	places := []string{
		"Place",
		"Caffee",
		"Shisha bar",
		"Kiosk",
	}

	/*
		BUILD STRUCTS
	*/

	locations := []location{}

	// Iterate over names
	for i := 0; i < len(names); i++ {

		// Iterate over places
		for j := 0; j < len(places); j++ {

			// Compose variables
			fullName := names[i] + " " + places[j]
			desc := fullName + "s description..."
			website := "www." + names[i] + places[j] + ".de"

			rand.Seed(time.Now().UnixNano())

			// Generate a random number 0 - 5
			price := rand.Intn(6)

			// Random keys generated by using google's encryption tools
			id := uuid.New().String()

			// Build struct from the variables
			tempLocation := location{
				name:    fullName,
				coordX:  0,
				coordY:  0,
				desc:    desc,
				website: website,
				price:   price,
				id:      id,
			}
			// Append to location array
			locations = append(locations, tempLocation)
		}
	}

	querries := ""

	// Write querries
	for _, location := range locations {

		// Setup a querry with 'printf' - like casting
		querries += fmt.Sprintf(addNewLocation, location.name, location.coordX, location.coordY, location.desc, location.website, location.price, location.id)

		// Add line break
		querries += "\n"
	}

	return querries, locations
}

// Generate comments about every location from every user
func generateComments(users []user, locations []location) string {
	querries := ""

	// Iterate over locations
	for _, location := range locations {

		// Iterate over users
		for _, user := range users {

			// Create variables
			title := user.name + "s comment..."

			// Lorem ipsum -> declared in ./querryTemplates.go
			content := lorem
			id := uuid.New().String()

			rand.Seed(time.Now().UnixNano())

			// Generate a random number 0 - 100
			ratingValue := rand.Intn(6)
			ratingID := uuid.New().String()

			// Generate querry for new comment
			querries += fmt.Sprintf(addNewComment, title, content, user.id, location.id, ratingID, id)

			querries += "\n"
			querries += fmt.Sprintf(addNewRating, ratingValue, location.id, ratingID)

			// Add line break
			querries += "\n"
		}
	}
	return querries
}

// Generate querries for features of a location
func generateFeatures(locations []location) string {

	querries := ""

	// generate a random boolean
	// p = probability
	randBool := func(p int) bool {
		// Generate random price
		// New seed of random numbers
		rand.Seed(time.Now().UnixNano())

		// Generate a random number 0 - 100
		n := rand.Intn(101)
		if n > p {
			return false
		}
		return true
	}

	// Iterate over locations
	for _, location := range locations {
		querries += fmt.Sprintf(declareFeatures, randBool(50), randBool(50), randBool(50), randBool(50), randBool(50), location.id)
		querries += "\n"
	}

	return querries
}
