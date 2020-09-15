package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.iosoftworks.com/EnterpriseNote/pkg/config"
	"go.uber.org/zap"
)

//Server this is the struct that contains the webserver information
type Server struct {
	config config.WebServerConfig
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
	zap.S().Info("Starting webserver...")
	server.config.Port = "8080"
	if server.config.Port == "" {
		server.config.Port = "8080"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(Notes)
}

//use mux to get us single notes
func (server Server) returnSingleNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, note := range Notes {
		if note.Id == key {
			json.NewEncoder(w).Encode(note)
		}
	}
}

func (server Server) createNewNote(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var note Note
	json.Unmarshal(reqBody, &note)
	note.Id = string(rune(len(Notes)))
	Notes = append(Notes, note)
	json.NewEncoder(w).Encode(note)
}

//run me to make the server work
func (server Server) handleRequests() {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/notes", server.returnAllNotes).Methods("GET")
	r.HandleFunc("/api/v1/note/{id}", server.returnSingleNote).Methods("GET")
	r.HandleFunc("/api/v1/note", server.createNewNote).Methods("POST")
	r.Handle("/", http.RedirectHandler("/web/", http.StatusPermanentRedirect))
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	log.Fatal(http.ListenAndServe(":8080", r))
}
