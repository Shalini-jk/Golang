package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	err := initDB(logger)

	if err != nil {
		fmt.Println("InitDB failed", err)
	}

	router := mux.NewRouter()

	// Use the global 'db' variable in your handlers
	router.HandleFunc("/userskills", getAllUserSkills).Methods("GET")
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", router)

}
