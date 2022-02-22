package main

import (
	"encoding/json"
	"io/ioutil"
)

type ma struct {
	index int
	clip  string
}

const fileLocation = "/.hidden/lists.json"

func errHanding(err error) {
	if err != nil {
		panic(err)
	}
}

func readJson() (elements []string) {
	fileBytes, err := ioutil.ReadFile(fileLocation)
	errHanding(err)
	err = json.Unmarshal(fileBytes, &elements)
	errHanding(err)
	return elements
}

func writeJson(element []string) {

	infoByte, err := json.Marshal(element)
	errHanding(err)
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
	errHanding(err)
}
