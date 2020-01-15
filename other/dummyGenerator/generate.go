package dummygenerator

import (
	"fmt"
	"kihmo"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

/*
	HERE ARE ALL GENERATE FUNCTIONS FOR EVERY TABLE

	EACH FUNCTION RETURNS BOTH THE QUERRIES AS STRING
	AND THE STRUCT SO THAT THE DATA USED TO WRITE THE
	QUERRY CAN BE REUSED
*/

// Generate user querries
func generateUsers() (string, []user) {

	// Array storing all names
	names := []string{
		"Louis",
		"Bartosz",
		"Jannik",
		"Kai",
		"Tim",
		"Bettina",
		"Aleksandra",
		"Jens",
		"Klaus",
		"Michael",
		"Rudolph",
		"Stefan",
		"Lukas",
		"Luca",
		"Adalbert",
	}

	// Array stroing all surnames
	surnames := []string{
		"Beul",
		"Burgiel",
		"Weyer",
		"Huntington",
		"Dell",
		"Jauch",
		"Schmidt",
		"Schneider",
		"Müller",
		"Spohr",
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
			password, err := kihmo.Hash(password, 60)
			if err != nil {
				panic(err)
			}

			// Random keys by using google's encrytpion tool
			comments := uuid.New().String()
			favlocation := uuid.New().String()
			ratings := uuid.New().String()
			id := uuid.New().String()

			// Build temporary user struct
			tempUser := user{
				name:        fullName,
				email:       email,
				password:    password,
				comments:    comments,
				favlocation: favlocation,
				ratings:     ratings,
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
		querries += fmt.Sprintf(addNewUser, user.name, user.password, user.email, user.comments, user.favlocation, user.ratings, user.id)

		// Add line break
		querries += "\n"
	}

	return querries, users
}

// Generate user querries
func generateEncodedUsers() (string, []user) {

	// Array storing all names
	names := []string{
		"Louis",
		"Bartosz",
		"Jannik",
		"Kai",
		"Tim",
		"Bettina",
		"Aleksandra",
		"Jens",
		"Klaus",
		"Michael",
		"Rudolph",
		"Stefan",
		"Lukas",
		"Luca",
		"Adalbert",
	}

	// Array stroing all surnames
	surnames := []string{
		"Beul",
		"Burgiel",
		"Weyer",
		"Huntington",
		"Dell",
		"Jauch",
		"Schmidt",
		"Schneider",
		"Müller",
		"Spohr",
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
			password, err := kihmo.Hash(password, 60)
			if err != nil {
				panic(err)
			}

			// Random keys by using google's encrytpion tool
			comments := uuid.New().String()
			favlocation := uuid.New().String()
			ratings := uuid.New().String()
			id := uuid.New().String()

			// Build temporary user struct
			tempUser := user{
				name:        fullName,
				email:       email,
				password:    password,
				comments:    comments,
				favlocation: favlocation,
				ratings:     ratings,
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
		querries += fmt.Sprintf(addNewUser, user.name, user.password, user.email, user.comments, user.favlocation, user.ratings, user.id)

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
		"Wirgeser",
		"Barteks",
		"Louis",
		"Lukas",
	}

	// Array storing place types (?) of the locaton
	places := []string{
		"Place",
		"Caffee",
		"Cafeteria",
		"Steak House",
		"Restaurant",
		"Mensa",
		"Bürgerhaus",
		"Dönerbude",
		"Bar",
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

			// Random keys generated by using google's encryption tools
			comments := uuid.New().String()
			users := uuid.New().String()
			ratings := uuid.New().String()
			id := uuid.New().String()

			// Build struct from the variables
			tempLocation := location{
				name:     fullName,
				coordX:   0,
				coordY:   0,
				desc:     desc,
				comments: comments,
				website:  website,
				users:    users,
				ratings:  ratings,
				id:       id,
			}

			// Append to location array
			locations = append(locations, tempLocation)
		}
	}

	querries := ""

	// Write querries
	for _, location := range locations {

		// Setup a querry with 'printf' - like casting
		querries += fmt.Sprintf(addNewLocation, location.name, location.coordX, location.coordY, location.desc, location.website, location.comments, location.users, location.ratings, location.id)

		// Add line break
		querries += "\n"
	}

	return querries, locations
}

// Genrate querries for the ratings table
// from existing users and location structs
// every user gives a rating to every location (0-10)
func generateRatings(users []user, locations []location) string {

	querries := ""

	// Iterate over locations
	for _, location := range locations {

		// Iterate over users
		for _, user := range users {

			// New seed of random numbers
			rand.Seed(time.Now().UnixNano())

			// Generate a random number for the rating 0 - 10
			rating := rand.Intn(11)

			// Generate querry
			querries += fmt.Sprintf(addNewRating, user.id, location.id, rating, location.ratings)

			// Add new line
			querries += "\n"
		}
	}

	return querries
}

// Generate querries assigning favorite users' locations
// Every user gets one random location as favorite
func generateFavLocations(users []user, locations []location) string {
	querries := ""

	// Iterate over users
	for _, user := range users {

		// Generate random index of the location
		// New seed of random numbers
		rand.Seed(time.Now().UnixNano())

		// Generate a random number for the rating 0 - 10
		locationIndex := rand.Intn(len(locations))

		querries += fmt.Sprintf(setFavoriteLocation, user.id, locations[locationIndex].id)

		// Add line break
		querries += "\n"
	}

	return querries
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

			// Generate querry
			querries += fmt.Sprintf(addNewComment, title, content, user.id, location.id, id)

			// Add line break
			querries += "\n"
		}
	}
	return querries
}

// Generate querries for features of a location
func generateFeatures(locations []location) string {

	querries := ""

	// Array of all features
	features := []string{
		"WiFi",
		"Etherner",
		"Power supply",
		"Pets",
		"Nice staff",
	}

	// Iterate over locations
	for _, location := range locations {

		// Iterate over features
		for _, feature := range features {

			// Create variables
			desc := "Very good " + feature
			id := uuid.New().String()

			// Generate querry
			querries += fmt.Sprintf(addNewFeature, feature, desc, location.id, id)

			// Add line break
			querries += "\n"
		}
	}

	return querries
}

// generate various dishes with a random price (1-5)
func generateDishes(locations []location) string {
	querries := ""

	// Temporary struct to make the querry generation easier
	type dish struct {
		name string
		typ  string
	}

	// Array of all dishes
	dishes := []dish{
		dish{
			"Coffee",
			"Beverage",
		},
		dish{
			"Tee",
			"Beverage",
		},
		dish{
			"Cheese cake",
			"Food",
		},
		dish{
			"Doughnut",
			"Food",
		},
		dish{
			"Sandwich",
			"Food",
		},
		dish{
			"Espresso",
			"Beverage",
		},
	}

	// Iterate over locations
	for _, location := range locations {

		// Iterate over dishes
		for _, dish := range dishes {

			// Generate random price
			// New seed of random numbers
			rand.Seed(time.Now().UnixNano())

			// Generate a random number for the rating 0 - 10
			price := rand.Intn(6) + 1

			id := uuid.New().String()

			// Generate querry
			querries += fmt.Sprintf(addNewDish, dish.name, dish.typ, price, location.id, id)

			// Add line break
			querries += "\n"
		}
	}

	return querries
}
