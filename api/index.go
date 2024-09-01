package handler

import (
	"json"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}