package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

type event struct {
	ID string `json:"ID"`
	Title  string `json:"Title"`
	Description string `json:"Description"`

}

type allEvent []event

var events = allEvent {
	{
		ID: "1",
		Title: "Introduction to Golang",
		Description: "Come join us for a change to learn how golang works and get to eventually try it out",

	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome home !")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	// convert r.body into a readable format

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Kindly update data for events" )
	}
	json.Unmarshal(reqBody, &newEvent)

	// Add the newly created event to the array forj events
	events = append(events, newEvent)

	// Return the 201 created status code
	w.WriteHeader(http.StatusCreated)

	// return the newly created events
	json.NewEncoder(w).Encode(newEvent)


}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8000", router))
}