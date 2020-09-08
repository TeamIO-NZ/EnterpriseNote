package web

import (
	"encoding/json"
	"fmt"
	"go.iosoftworks.com/EnterpriseNote/pkg/config"
	"go.uber.org/zap"
	"log"
	"net/http"
)

//Server this is the struct that contains the webserver information
type Server struct {
	config config.WebServerConfig
}

type Note struct {
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(Notes)
}

//run me to make the server work
func (server Server) handleRequests() {
	// api paths
	http.HandleFunc("/api/v1/notes", server.returnAllNotes)

	// frontend
	http.Handle(
		"/",
		http.FileServer(http.Dir("./site")),
	)

	zap.S().Info("Web server online")
	// ahhhh
	log.Fatal(http.ListenAndServe(":8080", nil))
}
