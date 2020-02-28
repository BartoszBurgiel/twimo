package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	// Sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// Repo represents a data base with its all methods
type Repo struct {
	db     *sql.DB
	path   string
	config config
}

// NewSQLite creates a new repository
func NewSQLite(path string) (*Repo, error) {
	r := &Repo{
		path: path,
	}
	r.config.init("../repository/config.json")
	fmt.Println(r.config)
	return r, r.initSQLite()
}

// initialize a repository
func (r *Repo) initSQLite() error {

	// Check if database folder exists
	// If not - mkdir
	p := path.Dir(r.path)
	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		fmt.Println(err)
		return err
	}

	// Open connection
	db, err := sql.Open("sqlite3", r.path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Read setup query from the file
	tempSetupDb, err := ioutil.ReadFile("../repository/queries/setup.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	setupDB := string(tempSetupDb)
	// Execute setup query
	_, error := db.Exec(setupDB)
	if error != nil {
		fmt.Println(err)
		panic(err)
	}

	// Add opened DB to repo struct
	r.db = db

	return nil
}
