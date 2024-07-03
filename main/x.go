package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux:=http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, Name: "John"},
			{ID: 2, Name: "bob"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})
	http.ListenAndServe(":8000", mux)

}
