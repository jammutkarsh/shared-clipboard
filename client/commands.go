package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const fileLocation = ".lists.json"

func readJson() (elements []string) {
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

func writeJson(element []string) {
	infoByte, err := json.Marshal(element)
	if err != nil {
		panic(err)
	}
	if len(readJson()) == 0 {
		// to make this a valid json file for appending
		clear()
	}
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
	if err != nil {
		panic(err)
	}
}

func add() {
	//	addCMD := flag.NewFlagSet("add", flag.ExitOnError)
	//	data := addCMD.String("c", "", "Enter your clip.")
	//	if len(os.Args) < 3 {
	//		fmt.Print(`Options:
	//	-c "your clip"
	//`)
	//		os.Exit(0)
	//	}
	//	err := addCMD.Parse(os.Args[2:])
	//	if err != nil {
	//		panic(err)
	//	}
	//	element := append(readJson(), *data)
	//	writeJson(element)
	POSTclip()
}

func clear() {
	clear := []byte("[]")
	_ = ioutil.WriteFile(fileLocation, clear, 0755)
	DELETEclips()
}

func deleteClip() {
	delCMD := flag.NewFlagSet("deleteClip", flag.ExitOnError)
	data := delCMD.Int("d", 0, "deleteClip a clip")
	if len(os.Args) < 3 {
		fmt.Print(`Options:	
	-d [value]
`)
		os.Exit(0)
	}
	err := delCMD.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}
	jsonObjects := readJson()
	var newJsonObjects []string
	for i, entry := range jsonObjects {
		if *data-1 == i {
			continue
		}
		newJsonObjects = append(newJsonObjects, entry)
	}
	infoByte, err := json.Marshal(newJsonObjects)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
}

func list() {
	GETclip()
	//readData := readJson()
	//for i, data := range readData {
	//	fmt.Printf("|%d|\t%s\t|\n", i+1, data)
	//}
}

func help() {
	fmt.Print(`Commands:	
	add    - to add clip to server
	deleteClip - to deleteClip a single data clip to the list.   
	list   - to display all clips.   
	clear  - to clear the list.  
`)
}
