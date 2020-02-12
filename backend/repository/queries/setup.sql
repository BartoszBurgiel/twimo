-- Table storing all user's information
-- name, password and email will be encoded with kihmo
-- Only neccesary information -> no user profile pages anyway
CREATE TABLE IF NOT EXISTS users (
	name 	    		TEXT,
	password	  		TEXT,
	email       	    TEXT,

	-- "xyz location is the fav location of n user"
	favlocation       	VARCHAR(16), 
	id					VARCHAR(16)
) ;

-- Stores all of the ratings 
CREATE TABLE IF NOT EXISTS ratings (
	value			NUMERIC,
    id            VARCHAR(16)
) ;

-- Stores all of the comments
CREATE TABLE IF NOT EXISTS comments (
	title 	   	VARCHAR(16),
	content  		TEXT,
	userID        VARCHAR(16),
	locationID	VARCHAR(16), 
	ratingID VARCHAR(16),
	id            VARCHAR(16)
) ;

-- Stores all location 
CREATE TABLE IF NOT EXISTS locations (
	name 	    		VARCHAR(16),
	coordX  		    NUMERIC,
	coordY            NUMERIC,
	descr              TEXT,
	webpage				VARCHAR(16),
	id				VARCHAR(16)
) ;

-- Features of each location
CREATE TABLE IF NOT EXISTS features (
	coffee BIT, 
	wifi  BIT,
	power BIT,
	music  BIT,
	food BIT,
	locationID VARCHAR(16)
) ;