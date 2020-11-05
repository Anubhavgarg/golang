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
	myRouter.HandleFunc("/createInstance", createInstance).Methods("POST")
	myRouter.HandleFunc("/deleteInstance/{instanceName}", deleteInstance).Methods("POST")
	myRouter.HandleFunc("/stopInstance/{instanceName}", stopInstance).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func deleteInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	url := "https://compute.googleapis.com/compute/beta/projects/western-notch-185412/zones/us-central1-a/instances/" + vars["instanceName"]
	method := "DELETE"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	var posted m
	_ = json.NewDecoder(res.Body).Decode(&posted)
	json.NewEncoder(w).Encode(posted)
}
func stopInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	url := "https://compute.googleapis.com/compute/beta/projects/western-notch-185412/zones/us-central1-a/instances/" + vars["instanceName"] + "/stop"
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	var posted m
	_ = json.NewDecoder(res.Body).Decode(&posted)
	json.NewEncoder(w).Encode(posted)
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
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	var posted m
	_ = json.NewDecoder(res.Body).Decode(&posted)
	json.NewEncoder(w).Encode(posted)
}

func main() {
	handleRequests()
}
