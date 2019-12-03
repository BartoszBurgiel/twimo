package sqlitedb

import "twimo/backend/core"

func (r Repo) GetUser(userID string) *core.User {
	return nil
}

func (r Repo) GetComment(commentID string) *core.Comment {
	return nil
}

func (r Repo) GetLocation(locationID string) *core.Location {
	return nil
}

func (r Repo) GetCommentsOfLocation(locationID string) []*core.Comment {
	return nil
}

func (r Repo) GetLocationsFavUsers(locationID string) []*core.User {

	return nil
}

func (r Repo) GetLocationRatings(locationID string) []*core.Rating {

	return nil
}
