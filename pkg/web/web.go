package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Server this is the struct that contains the webserver information
type Server struct {
	port int
}

type Note struct {
	Id      string `json:"Id"`
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
		Note{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Note{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
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

//use mux to get us single notes
func (server Server) returnSingleNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, note := range Notes {
		if note.Id == key {
			json.NewEncoder(w).Encode(note)
		}
	}
}

//run me to make the server work
func (server Server) handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", server.homePage)
	myRouter.HandleFunc("/notes", server.returnAllNotes)
	myRouter.HandleFunc("/note/{id}", server.returnSingleNote)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(server.port), myRouter))
}
