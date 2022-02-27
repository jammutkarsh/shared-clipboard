package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// File Check
	if !FileExists(fileLocation) {
		_, err := os.Create(fileLocation)
		clear()
		if err != nil {
			panic(err)
		}
	}
	// Init router
	router := mux.NewRouter()

	// Route handles & endpoints
	router.HandleFunc("/fetch", getClips).Methods("GET")
	router.HandleFunc("/add", addClip).Methods("POST")
	router.HandleFunc("/delete/{index}", deleteClip).Methods("DELETE")
	router.HandleFunc("/clear", clearClips).Methods("DELETE")

	// Start server
	fmt.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
