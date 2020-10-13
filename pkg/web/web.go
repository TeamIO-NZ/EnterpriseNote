package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.iosoftworks.com/EnterpriseNote/pkg/config"
	"go.uber.org/zap"
)

//Server this is the struct that contains the webserver information
type Server struct {
	config config.WebServerConfig
}

//Note a note object for json
type Note struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

//Notes a note array
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

//------------------------------JSON Webrequests Hander functions--------------------------------//

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
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to return
	key := vars["id"]

	for _, note := range Notes {
		if note.Id == key {
			json.NewEncoder(w).Encode(note)
			return
		}
	}
}

func (server Server) createNewNote(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var note Note
	json.Unmarshal(reqBody, &note)
	Notes = append(Notes, note)
	json.NewEncoder(w).Encode(note)
}

func (server Server) deleteNote(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	for index, note := range Notes {
		if note.Id == id {
			Notes = append(Notes[:index], Notes[index+1:]...)
			return
		}
	}
}
func (server Server) updateNote(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	for index, note := range Notes {
		if note.Id == vars["id"] {
			Notes = append(Notes[:index], Notes[index+1:]...)

			var note Note
			_ = json.NewDecoder(r.Body).Decode(&note)
			note.Id = vars["id"]
			Notes = append(Notes, note)
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	json.NewEncoder(w).Encode(Notes)

}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
func createTable() {
	db := createConnection()
	defer db.Close()
	sqlStatement := `CREATE TABLE IF NOT EXISTS notes (
		userid SERIAL PRIMARY KEY,
		title TEXT,
		description TEXT,
		contents TEXT
	);`
	res, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("%s\n ", res)
}

//------------------------------SQL Hander functions--------------------------------//

//run me to make the server work
func (server Server) handleRequests() {

	createTable()
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/notes", server.returnAllNotes).Methods("GET")
	r.HandleFunc("/api/v1/note/{id}", server.returnSingleNote).Methods("GET")
	r.HandleFunc("/api/v1/note", server.createNewNote).Methods("POST")
	r.HandleFunc("/api/v1/note/{id}", server.updateNote).Methods("PUT")

	r.HandleFunc("/api/v1/note/{id}", server.deleteNote).Methods("DELETE")
	r.Handle("/", http.RedirectHandler("/web/", http.StatusPermanentRedirect))
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	log.Fatal(http.ListenAndServe(":8080", r))
}
