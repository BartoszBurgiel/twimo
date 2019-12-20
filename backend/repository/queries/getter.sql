-- Get location
SELECT name, coordX, coordY, desc, comments, users, ratings, id
FROM locations 
WHERE id = ? ; 

-- get locations comment
SELECT title, content, userID, locationID, id
FROM comments 
WHERE locationID = ? ; 

-- get locations fav users 
SELECT name, email, id 
FROM users, location_fav_users 
WHERE users.id = location_fav_users.userID 
AND location_fav_users.locationID = "64201" ;

-- get user 
SELECT name, password, email, comments, favlocation, ratings, id 
FROM users 
WHERE id = ? ; 

-- get comment 
SELECT title, content, userID, locationID, id
FROM comments 
WHERE id = ? ; 

-- get location ratings
SELECT userID, value, id
FROM ratings
WHERE locationID = ? ; 