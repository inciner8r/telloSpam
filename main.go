package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func send(link string, responseBody *bytes.Buffer) {
	resp, err := http.Post(link, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("errrror :", err)
	}
	sb := string(body)
	log.Println(sb)
	if len(sb) > 0 {
		log.Println("error hui", len(sb))
		time.Sleep(3 * time.Second)
		http.Post(link, "application/json", responseBody)
	}
	time.Sleep(2 * time.Second)
}
func main() {
	type postReq struct {
		IsInstagramInAppBrowser bool   `json:"isInstagramInAppBrowser"`
		IsSenderRevealed        bool   `json:"isSenderRevealed"`
		Tell                    string `json:"tell"`
		UserId                  int    `json:"userId"`
		Limit                   int    `json:"limit"`
	}
	link := "https://api.tellonym.me/tells/new"

	for i := 0; i < 99; i++ {
		message := "testing spam lulul final" + strconv.Itoa(i)

		postBody, err := json.Marshal(postReq{IsInstagramInAppBrowser: false, IsSenderRevealed: false, Tell: message, UserId: 82491097, Limit: 99})
		if err != nil {
			log.Fatalf("An Error Occured in marshal. %v", err)
		}
		responseBody := bytes.NewBuffer(postBody)
		send(link, responseBody)
	}
}
