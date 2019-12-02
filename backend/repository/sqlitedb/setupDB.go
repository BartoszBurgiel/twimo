package sqlitedb

const setupDB = `CREATE TABLE IF NOT EXISTS 'users' (
	'name' 	    		VARCHAR(16),
	'password'  		VARCHAR(16),
	'email'             VARCHAR(16),
	'comments'		    VARCHAR(16), 
	'favorite-location' VARCHAR(16), 
	'ratings'           VARCHAR(16)
	) ;

CREATE TABLE IF NOT EXISTS 'comments' (
	'title' 	   	VARCHAR(16),
	'content'  		TEXT,
	'user'          VARCHAR(16),
	'location'		VARCHAR(16), 
	'id'            VARCHAR(16)
	) ;

CREATE TABLE IF NOT EXISTS 'locations' (
	'name' 	    		VARCHAR(16),
	'coordX'  		    NUMERIC,
	'coordY'            NUMERIC,
	'desc'              TEXT,
	'comments'		    VARCHAR(16), 
	'users'             VARCHAR(16), 
	'ratings'           VARCHAR(16)
	) ;
`
