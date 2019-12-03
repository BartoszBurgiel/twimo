package main

import (
	"fmt"
	"net/http"
	"twimo/backend/repository/sqlitedb"
	"twimo/backend/server"
)

func main() {

	repo, err := sqlitedb.New("../repository/sqlitedb/sql/repo.db")
	if err != nil {
		fmt.Println(err)
	}

	s, err := server.NewServer(repo)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", s)
}
