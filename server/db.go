package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const FILELOCATION = "server/.list.json"

func readJSON() (clips []string) {
	if fileExist() {
		fileBytes, err := ioutil.ReadFile(FILELOCATION)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(fileBytes, &clips)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("File can not created")
	}
	return clips
}

func writeJSON(data string) {
	element := append(readJSON(), data)
	infoByte, err := json.Marshal(element)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(FILELOCATION, infoByte, 0644)
	if err != nil {
		panic(err)
	}
}
func fileExist() bool {
	_, err := os.Stat(FILELOCATION)

	if os.IsNotExist(err) {
		//creating file
		cleanDB()
		return false
		//file not exist
	}
	return true
}

func cleanDB() {
	err := os.WriteFile(FILELOCATION, []byte("[]"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
