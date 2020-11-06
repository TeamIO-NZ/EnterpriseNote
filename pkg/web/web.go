package web

import (
	"database/sql"
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable

	"github.com/gorilla/mux"   // used to get the params from the route
	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      //sql driver. blank is required
	"go.iosoftworks.com/EnterpriseNote/pkg/config"
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
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

//Notes a note array
var Notes []models.Note
var port = "8082"

//Start this is the function that starts the webserver
func (server Server) Start() {
	zap.S().Info("Starting webserver...")
	zap.S().Info("Go to http://localhost:8082/web/#/")
	server.config.Port = port
	if server.config.Port == "" {
		server.config.Port = port
		//zap.S().Warn("No webserver port config detected, using 8080.")
	}
	// Notes = []models.Note{
	// 	models.Note{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	// 	models.Note{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	// }

	server.HandleRequests()

}

//HandleRequests run me to make the server work
func (server Server) HandleRequests() {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/notes", server.ReturnAllNotes).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.ReturnSingleNote).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/note", server.CreateNewNote).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.UpdateNote).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.DeleteNote).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/users", server.ReturnAllUsers).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.ReturnSingleUser).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user", server.CreateNewUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.UpdateUser).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.DeleteUser).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/login", nil).Methods("GET")

	r.Handle("/", http.RedirectHandler("/web/", http.StatusPermanentRedirect)).Methods("GET", "OPTIONS")
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	log.Fatal(http.ListenAndServe(":"+port, r))
}

// func (server Server) homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the home page!")
// 	fmt.Println("Endpoint Hit: homepage")
// }

//------------------------------SQL Hander functions--------------------------------//
// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file

	err := godotenv.Load(".env")

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	// check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

//TODO function for dropping and rebuilding tables.
//TODO function for loading in a csv for a temp table or something
func createTable() {
	//creates database connection
	db := createConnection()
	//prepares to close database when done
	defer db.Close()
	//create the base notes table for if it doesn't exist
	sqlStatement := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		password TEXT,
		email TEXT
	);`

	//execute the sql statement and return a response
	res, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	//print the response maybe
	fmt.Printf("%s\n ", res)

	//TODO rebuild this table function
	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS notes (
			id SERIAL PRIMARY KEY,
			name TEXT,
			password TEXT,
			email TEXT
		);`
	//execute the sql statement and return a response
	res, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	//print the response maybe
	fmt.Printf("%s\n ", res)
}
