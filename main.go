package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

type m map[string]interface{}

// Existing code from above
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/createInstance", createInstance).Methods("GET")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createInstance(w http.ResponseWriter, r *http.Request) {
	url := "https://content-compute.googleapis.com/compute/beta/projects/western-notch-185412/zones/us-central1-a/instances"
	method := "POST"
	// decoder := json.NewDecoder(r.Body)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	payload := strings.NewReader(string(reqBody))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authorization", "Bearer ya29.A0AfH6SMDCpmgdwaYjgG5xgUSmI0tx1S5_rAHsdeawFJpRuOV7rsupflvGoRQ8ZomsafH1noC8urlrRVdaeC5hhf5g9JIGWkqeydm7P67xKkJPyQ-RmANq1fQd-9d4NZT1tgmicmDCgJxyGdm8DrguHcUUgXAxPNDXyn7oAmlTPhgMwPTl0yS7")
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	var posted m
	_ = json.NewDecoder(res.Body).Decode(&posted)
	json.NewEncoder(w).Encode(posted)
}

func main() {
	handleRequests()
}
