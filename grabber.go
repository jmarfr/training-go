package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JsonResp struct {
	Args    map[string]interface{} `json:"args"`
	Headers RespHeaders            `json:"headers"`
	Origin  string                 `json:"origin"`
	Url     string                 `json:"url"`
}

type RespHeaders struct {
	Accept         string `json:"Accept"`
	AcceptEncoding string `json:"Accept-Encoding"`
	Host           string `json:"Host"`
	UserAgent      string `json:"User-Agent"`
}

func main() {

	// build the request
	req, err := http.NewRequest("GET", "http://httpbin.org/get?pouetpouet=toto", nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Set("accept", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	var record JsonResp

	// if resp.StatusCode == http.StatusOK {
	// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	bodyString := string(bodyBytes)
	// 	fmt.Println(bodyString)
	// }

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println("Decoder error: ", err)
	}

	fmt.Println("Records: ", record)
	fmt.Println("Origin: ", record.Origin)
	fmt.Println("Arg1: ", record.Args["pouetpouet"])
}
