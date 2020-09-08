package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//Server this is the struct that contains the webserver information
type Server struct {
	port int
}

//Start this is the function that starts the webserver
func (server Server) Start() {
	if server.port == 0 {
		server.port = 8080
		//zap.S().Warn("No webserver port config detected, using 8080.")
	}
	server.handleRequests()

}
func (server Server) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homepage")
}

//run me to make the server work
func (server Server) handleRequests() {
	http.HandleFunc("/", server.homePage)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(server.port), nil))
}
