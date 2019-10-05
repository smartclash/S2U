package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/smartclash/S2U/database"
	"github.com/smartclash/S2U/handlers"
)

func main() {
	fmt.Println("Starting Go server for Sub2Unlock")

	r := mux.NewRouter()
	registerRoutes(r)

	var err error
	database.Database, err = gorm.Open("sqlite3", "test.sqlite")
	if err != nil {
		panic("Could not connect to a database")
	}

	defer func() {
		err = database.Database.Close()
		if err != nil {
			fmt.Println("Could not close the database connection successfully")
		}
	}()
	database.RunMigrations()

	if err = http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error binding to port 8080. Could not resolve connections")
	}
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.Index)
}
