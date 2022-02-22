package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

//func main() {
//	//signIn()
//}

func signIn() {
	var user string
	fmt.Print("Enter username: ")
	_, _ = fmt.Scan(&user)
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
	password := string(bytePassword)
	userLocation := "client/.user/" + user
	fmt.Println(userLocation)
	fmt.Println(password)
	if fileCheck(userLocation) {
		fmt.Println("user already exists")
	} else {
		fmt.Println("Account created")
		_, _ = os.Create(userLocation)
	}
}

func fileCheck(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
