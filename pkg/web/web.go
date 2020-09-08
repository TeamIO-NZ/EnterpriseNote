package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//Server this is the struct that contains the webserver information
type Server struct {
	port int
}

type Note struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var Notes []Note

//Start this is the function that starts the webserver
func (server Server) Start() {
	if server.port == 0 {
		server.port = 8080
		//zap.S().Warn("No webserver port config detected, using 8080.")
	}
	Notes = []Note{
		Note{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Note{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	server.handleRequests()

}
func (server Server) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homepage")
}
func (server Server) returnAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndpointHit: return all notes")
	json.NewEncoder(w).Encode(Notes)
}

//run me to make the server work
func (server Server) handleRequests() {
	http.HandleFunc("/", server.homePage)
	http.HandleFunc("/notes", server.returnAllNotes)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(server.port), nil))
}
