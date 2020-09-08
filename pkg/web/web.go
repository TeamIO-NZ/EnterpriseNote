package web

import (
	"fmt"
	"go.iosoftworks.com/EnterpriseNote/pkg/config"
	"log"
	"net/http"
)

//Server this is the struct that contains the webserver information
type Server struct {
	config config.WebServerConfig
}

//Start this is the function that starts the webserver
func (server Server) Start() {
	if server.config.Port == "" {
		server.config.Port = "8080"
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
	http.Handle(
		"/",
		http.FileServer(http.Dir("./site")),
	)
	log.Fatal(http.ListenAndServe(":"+server.config.Port, nil))
}
