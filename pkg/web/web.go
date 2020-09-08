package web

import (
	"fmt"
	"log"
	"net/http"
)

type WebServer struct {
	port int
}

func (server WebServer) Start() {
	if server.port == 0 {
		server.port = 8080
	}
	handleRequests()

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homepage")
}

//run me to make the server work
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
