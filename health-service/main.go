package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnHealth")
	//json.NewEncoder(w).Encode(Articles)
	fmt.Fprintf(w, "{\"status\":\"OK\"}")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: home")
	//json.NewEncoder(w).Encode(Articles)
	fmt.Fprintf(w, "{\"hello\":\"world\"}")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homeHandler)
	myRouter.HandleFunc("/health", returnHealth)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
