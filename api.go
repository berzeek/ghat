package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the api!")
	fmt.Println("Endpoint Hit: homePage")
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	content, err := ioutil.ReadFile("ghat_log.txt")

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(string(content)))
	return
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/messages", getMessages)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
