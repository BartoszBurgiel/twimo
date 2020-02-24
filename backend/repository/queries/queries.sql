-- INSERT INTO:

-- set (as in change) user's favorite location
UPDATE users SET users.favlocation = ? WHERE users.id = ?

-- modify location's features
UPDATE features SET features.coffee = ? WHERE features.locationID = ?
UPDATE features SET features.wifi = ? WHERE features.locationID = ?
UPDATE features SET features.power = ? WHERE features.locationID = ?
UPDATE features SET features.music = ? WHERE features.locationID = ?
UPDATE features SET features.food = ? WHERE features.locationID = ?


-- SELECT:

-- get location avr. rating
SELECT AVG(value) FROM ratings WHERE ratings.locationID = ?

-- get top 20 locations sorted after:

  -- distance from user (LATER)
  
  -- rating
  SELECT locations.name, locations.descr, locations.id, locations.price, AVG(ratings.value)
  FROM locations 
  LEFT JOIN ratings
  ON locations.id = ratings.locationID
  ORDER BY AVG(ratings.value) ASC
  LIMIT 20

  -- number of fav users
  SELECT locations.name, locations.descr, locations.id, locations.price, AVG(ratings.value)
  FROM locations 
  LEFT JOIN users
  ON locations.id = users.favlocation
  ORDER BY COUNT(users.favlocation) ASC
  LIMIT 20

  -- price
  SELECT TOP 20 locations.name, locations.descr, locations.id, locations.price, AVG(ratings.value)
  FROM locations 
  ORDER BY locations.price ASC

-- DELETE:

-- user
DELETE FROM users WHERE id=?

-- location
DELETE FROM locations WHERE id=?

-- rating
DELETE FROM ratings WHERE id = ?
