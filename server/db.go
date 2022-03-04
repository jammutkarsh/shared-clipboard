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
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &elements)
	if err != nil {
		panic(err)
	}
	return elements
}

func writeJson(element []Clipboard) {
	infoByte, err := json.Marshal(element)
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
}

//func errHanding(err error) {
//	if err != nil {
//		return
//	}
//}
