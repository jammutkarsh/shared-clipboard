package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Clipboard struct {
	Clip string `json:"clip"`
}

const fileLocation = ".lists.json"

func readJson() (elements []Clipboard) {
	fileBytes, err := ioutil.ReadFile(fileLocation)
	errHanding(err)
	err = json.Unmarshal(fileBytes, &elements)
	errHanding(err)
	return elements
}

func writeJson(element []Clipboard) {
	infoByte, err := json.Marshal(element)
	errHanding(err)
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
	errHanding(err)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func clear() {
	clear := []byte("[]")
	err := ioutil.WriteFile(fileLocation, clear, 0755)
	errHanding(err)
}

func errHanding(err error) {
	if err != nil {
		return
	}
}
