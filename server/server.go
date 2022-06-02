package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/fetchClips", getClips).Methods("GET")
	router.HandleFunc("/addClips", postClip).Methods("POST")
	router.HandleFunc("/clear", clear).Methods("DELETE")
	fileExist()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
