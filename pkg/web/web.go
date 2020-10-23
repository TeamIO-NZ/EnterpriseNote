package web

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //sql driver. blank is required
	"go.iosoftworks.com/EnterpriseNote/pkg/config"
	"go.uber.org/zap"
)

//Server this is the struct that contains the webserver information
type Server struct {
	config config.WebServerConfig
}

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//Note a note object for json
type Note struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

//Notes a note array
var Notes []Note
var port = "8080"

//Start this is the function that starts the webserver
func (server Server) Start() {
	zap.S().Info("Starting webserver...")
	server.config.Port = port
	if server.config.Port == "" {
		server.config.Port = port
		//zap.S().Warn("No webserver port config detected, using 8080.")
	}
	Notes = []Note{
		Note{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Note{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	server.HandleRequests()

}

//HandleRequests run me to make the server work
func (server Server) HandleRequests() {

	createTable()
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/notes", server.ReturnAllNotes).Methods("GET")
	r.HandleFunc("/api/v1/note/{id}", server.ReturnSingleNote).Methods("GET")
	r.HandleFunc("/api/v1/note", server.CreateNewNote).Methods("POST")
	r.HandleFunc("/api/v1/note/{id}", server.UpdateNote).Methods("PUT")

	r.HandleFunc("/api/v1/note/{id}", server.DeleteNote).Methods("DELETE")
	r.Handle("/", http.RedirectHandler("/web/", http.StatusPermanentRedirect))
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	log.Fatal(http.ListenAndServe(":"+port, r))
}
func (server Server) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homepage")
}

//------------------------------SQL Hander functions--------------------------------//

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

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
		id SERIAL PRIMARY KEY,
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
