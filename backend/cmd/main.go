package main

import (
	"fmt"
	"net/http"
	"twimo/backend/repository"
	"twimo/backend/server"
)

func main() {
	fmt.Println("Server started")
	fmt.Println("Initializing the database")
	repo, err := repository.NewSQLite("../repository/sqlitedb/sql/repo.db")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Launching the server")
	s, err := server.NewServer(repo)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", s)
}
