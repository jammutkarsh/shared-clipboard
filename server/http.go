package main

import (
	"encoding/json"
	"net/http"
)

func getClips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(readJSON())
	if err != nil {
		return
	}
}

func postClip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newClip string
	err := json.NewDecoder(r.Body).Decode(&newClip)
	writeJSON(newClip)
	err = json.NewEncoder(w).Encode(&newClip)
	if err != nil {
		return
	}
}

func clear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cleanDB()
	// clearing db by overwriting a new file on it.
}
