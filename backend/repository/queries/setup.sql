-- Table storing all user's information
-- name, password and email will be encoded with kihmo
CREATE TABLE IF NOT EXISTS 'users' (
	'name' 	    		TEXT,
	'password'  		TEXT,
	'email'             TEXT,
	'comments'		    VARCHAR(16), 
	'favlocation'       VARCHAR(16), 
	'ratings'           VARCHAR(16), 
	'id'				VARCHAR(16)
) ;

-- Stores all of the ratings 
CREATE TABLE IF NOT EXISTS 'ratings' (
	'userID' 		VARCHAR(16), 
	'locationID' 	VARCHAR(16), 
	'value'			NUMERIC,
    'id'            VARCHAR(16)
) ;

-- Stores all of the comments
CREATE TABLE IF NOT EXISTS 'comments' (
	'title' 	   	VARCHAR(16),
	'content'  		TEXT,
	'userID'        VARCHAR(16),
	'locationID'	VARCHAR(16), 
	'id'            VARCHAR(16)
) ;


-- Table that assigns one favorite location to 
-- a user 
CREATE TABLE IF NOT EXISTS 'location_fav_users' (
	'userID' 		VARCHAR(16), 
	'locationID' 	VARCHAR(16)
) ;

-- Stores all features of a location (electricity, internet, ...)
CREATE TABLE IF NOT EXISTS 'location_features' (
	'name'			VARCHAR(16), 
	'desc'			VARCHAR(16),
	'locationID'	VARCHAR(16),
	'featuresID'	VARCHAR(16)
)

-- Stores dishes for the thumbnails 
-- Multiple dishes can be stored here, but only one will be chosen for the thumbnail later on 
CREATE TABLE IF NOT EXISTS 'dishes' (
	'name' 			VARCHAR(16),
	'type'			VARCHAR(16), 
	'price'			NUMERIC,
	'locationID'	VARCHAR(16),
	'ID' 	    	VARCHAR(16) 
) ;

-- Stores all location 
CREATE TABLE IF NOT EXISTS 'locations' (
	'name' 	    		VARCHAR(16),
	'coordX'  		    NUMERIC,
	'coordY'            NUMERIC,
	'desc'              TEXT,
	'comments'		    VARCHAR(16), 
	'users'             VARCHAR(16), 
	'ratings'           VARCHAR(16), 
	'id'				VARCHAR(16)
) ;


