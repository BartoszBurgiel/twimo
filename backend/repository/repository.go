package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	// Sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// Repo represents a data base with its all methods
type Repo struct {
	db   *sql.DB
	path string
}

// NewSQLite creates a new repository
func NewSQLite(path string) (*Repo, error) {
	r := &Repo{
		path: path,
	}

	return r, r.initSQLite()
}

// initialize a repository
func (r *Repo) initSQLite() error {

	// Check if database folder exists
	// If not - mkdir
	p := path.Dir(r.path)
	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		return err
	}

	// Open connection
	db, err := sql.Open("sqlite3", r.path)
	if err != nil {
		return err
	}

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
