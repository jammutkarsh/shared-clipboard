package main

import (
	"fmt"
	"io/ioutil"
)

const fileLocation = ".lists.json"

func add() {
	POSTclip()
}

func clear() {
	clear := []byte("[]")
	_ = ioutil.WriteFile(fileLocation, clear, 0755)
	DELETEclips()
}

func deleteClip() {
	// call single delete http method.
}

func list() {
	GETclip()
}

func help() {
	fmt.Print(`Commands:	
	add    - to add clip to server
	deleteClip - to deleteClip a single data clip to the list.   
	list   - to display all clips.   
	clear  - to clear the list.  
`)
}

// TODO: print clips in a well formatted manner.
