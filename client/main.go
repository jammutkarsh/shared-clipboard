package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no commands provided")
		os.Exit(1)
	}
	cmd := os.Args[1]

	switch cmd {
	case "add":
		add()
	case "list":
		list()
	case "clear":
		clear()
	case "deleteClip":
		deleteClip()
	case "help":
		help()
	default:
		help()
	}
}
