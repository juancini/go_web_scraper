package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []string{"John", "Mark", "Tom", "Jerry"}

	// convert to user list to JSON
	usersJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set content type
	w.Header().Set("Content-Type", "application/json")

	// write the JSON response to the client
	w.Write(usersJson)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World!")
	})

	router.HandleFunc("/users", getUsersHandler).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)
}
