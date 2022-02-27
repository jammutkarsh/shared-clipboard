package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Clipboard struct {
	Index int    `json:"index"`
	Clip  string `json:"clip"`
}

const fileLocation = "/Users/utkarshchourasia/code/go/go-projects/shared-clipboard/server/.lists.json"

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
	_ = ioutil.WriteFile(fileLocation, clear, 0755)
}

func errHanding(err error) {
	if err != nil {
		panic(err)
	}
}
