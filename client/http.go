package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "http://localhost:8080"

func GETclip() {
	resp, err := http.Get(URL + "/fetch")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	//resp, err := http.Get("http://localhost:8080/fetch")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	////We Read the response body on the line below.
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	////Convert the body to type string
	////sb := string(body)
	////fmt.Println(sb)
	//fmt.Println(resp)
}

func POSTclip() {
	//text := "clip" + string(rand.Int())
	//data := url.Values{
	//	//"index": {string(rune(rand.Int()))},
	//	"clip": {"Hello"},
	//}
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"clip": "Clip 1",
		//"email": "Toby@example.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(URL+"/add", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
func DELETEclips() {
	//fmt.Println("4. Performing Http Delete...")
	//todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal("[]")
	req, err := http.NewRequest(http.MethodDelete, URL+"/clear", bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

// TODO: respond back to use for bad response from client side.
