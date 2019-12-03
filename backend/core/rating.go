package core

// Rating represents a row in the ratings table
type Rating struct {
	UserID     string
	LocationID string
	Value      int
}
