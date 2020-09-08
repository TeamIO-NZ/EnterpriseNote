package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type WebServer struct {
	port int
}

func (server WebServer) Start() {
	if server.port == 0 {
		server.port = 8080
		//zap.S().Warn("No webserver port config detected, using 8080.")
	}
	server.handleRequests()

}
func (server WebServer) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homepage")
}

//run me to make the server work
func (server WebServer) handleRequests() {
	http.HandleFunc("/", server.homePage)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(server.port), nil))
}
