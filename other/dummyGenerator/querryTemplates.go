package dummygenerator

// SQL querry to add one new user
const addNewUser = `INSERT INTO users VALUES('%s', '%s', '%s', '%s', '%s') ;`

// SQL querry to add one new location
const addNewLocation = `INSERT INTO locations VALUES('%s', '%d', '%d', '%s', '%s', '%d', '%s') ;`

// SQL querry to add one new rating
const addNewRating = `INSERT INTO ratings VALUES('%d', '%s','%s') ;`

// SQL querry to add one new comment
const addNewComment = `INSERT INTO comments VALUES('%s', '%s', '%s', '%s', '%s', '%s') ;`

// SQL querry to generate features
const declareFeatures = `INSERT INTO features VALUES('%t', '%t', '%t', '%t', '%t', '%s') ;`

const lorem = `Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.`
