package web

import (
	"database/sql"
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable
	"strings"

	"github.com/gorilla/handlers"
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
	db     *sql.DB
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
	var nukeDatabase bool
	server.db, nukeDatabase = CreateConnection()
	server.HandleRequests()
	server.db.Close()
	if nukeDatabase {
		CreateTable(server.db)
	} else {
		firstRun(server.db)
	}

}

//HandleRequests run me to make the server work
func (server Server) HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/notes", server.ReturnAllNotes).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.ReturnSingleNote).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/note", server.CreateNewNote).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}/{userid}", server.UpdateNote).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.DeleteNote).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/users", server.ReturnAllUsers).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.ReturnSingleUser).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user/name/{username}", server.ReturnSingleUserByName).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user/email/{email}", server.ReturnSingleUserByEmail).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user", server.CreateNewUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.UpdateUser).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.DeleteUser).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/usersettings/{id}", server.ReturnSingleUserSettings).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/usersettings", server.CreateNewUserSettings).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/usersettings/{id}", server.UpdateUserSettings).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/usersettings/{id}", server.DeleteUserSettings).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/usersnotes/{id}", server.GetAllNotesUserHasAccessTo).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/login/{username}/{password}", server.Login).Methods("GET", "OPTIONS")

	r.HandleFunc("/api/v1/", server.CheckConnection).Methods("GET", "OPTIONS")

	r.Handle("/", http.RedirectHandler("/web/", http.StatusPermanentRedirect))
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "Origin", "Accept", "Access-Control-Request-Method", "Access-Control-Allow-Methods"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(r)))
}

//------------------------------SQL Hander functions--------------------------------//

//CreateConnection create connection with postgres db
func CreateConnection() (*sql.DB, bool) {
	err := godotenv.Load(".env") // load .env file
	var postgresString string    // store the string
	nukeDatabase := false
	if strings.ToLower(os.Getenv("USE_LOCAL_INSTANCE")) == "true" { // use local postgres install
		postgresString = os.Getenv("LOCAL_POSTGRES_URL")           // yes do
		if strings.ToLower(os.Getenv("NUKE_DATABASE")) == "true" { // nuke database
			nukeDatabase = true
		}
	} else {
		postgresString = os.Getenv("ALT_POSTGRES_URL")
	}
	db, err := sql.Open("postgres", postgresString) // Open the connection
	if err != nil {
		panic(err)
	}
	err = db.Ping() // check the connection
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db, nukeDatabase // return the connection
}

func firstRun(db *sql.DB) {
	PingOrPanic(db)
	//prepares to close database when done
	defer db.Close()
	//create the base notes table for if it doesn't exist
	sqlStatement := `CREATE TABLE IF NOT EXISTS userSettings (
		id Serial PRIMARY KEY,
		viewers integer[],
		editors integer[]
	);`
	Execute(db, sqlStatement)
	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS users (
		userId serial PRIMARY KEY,
		name TEXT,
		password TEXT,
		gender TEXT,
		email TEXT,
		token TEXT,
		userSettingsId int default 1
	);`
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS notes (
			id serial PRIMARY KEY,
			title TEXT,
			description TEXT,
			contents TEXT,
			owner INT,
			viewers integer[],
			editors integer[],
			FOREIGN KEY (owner)	REFERENCES users (userId) on delete cascade on update cascade
		);`
	Execute(db, sqlStatement)
}
