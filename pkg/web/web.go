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
	"github.com/lib/pq"
	_ "github.com/lib/pq" //sql driver. blank is required
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
		createTable(server.db)
	}
}

//HandleRequests run me to make the server work
func (server Server) HandleRequests() {

	//createTable(server.db)
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

// func (server Server) homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the home page!")
// 	fmt.Println("Endpoint Hit: homepage")
// }

//------------------------------SQL Hander functions--------------------------------//

//CreateConnection create connection with postgres db
func CreateConnection() (*sql.DB, bool) {
	// load .env file

	err := godotenv.Load(".env")
	var postgresString string
	nukeDatabase := false
	if strings.ToLower(os.Getenv("USE_LOCAL_INSTANCE")) == "true" {
		postgresString = os.Getenv("LOCAL_POSTGRES_URL")
		if strings.ToLower(os.Getenv("NUKE_DATABASE")) == "true" {
			nukeDatabase = true
		}
	} else {
		postgresString = os.Getenv("ALT_POSTGRES_URL")
	}
	// Open the connection
	db, err := sql.Open("postgres", postgresString)
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
	return db, nukeDatabase
}

//populates database with fake data
func createTable(db *sql.DB) {
	PingOrPanic(db)
	//prepares to close database when done
	defer db.Close()

	//create the base notes table for if it doesn't exist
	sqlStatement := `DROP TABLE IF EXISTS notes;`
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP TABLE IF EXISTS users;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS user_id_sequence;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS notes_userid_sequence;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS usersettings_id_sequence;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS zero_index_auto_increment;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)
	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP TABLE IF EXISTS userSettings;`
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS userSettings (
			id Serial PRIMARY KEY,
			viewers integer[],
			editors integer[]
		);`
	Execute(db, sqlStatement)
	UserSettings := []models.UserSettings{
		{
			ID:      1,
			Viewers: []int{6, 2, 3},
			Editors: []int{4, 5},
		},
		{
			ID:      2,
			Viewers: []int{1, 2, 3},
			Editors: []int{4, 5},
		},
		{
			ID:      3,
			Viewers: []int{1, 2, 3},
			Editors: []int{4, 5},
		},
	}
	for _, userSettings := range UserSettings {
		fmt.Println(userSettings.ID)
		var id int64
		sqlStatement := `INSERT INTO userSettings (id, viewers,editors) VALUES ($1,$2, $3) RETURNING id`

		err := db.QueryRow(sqlStatement, userSettings.ID, pq.Array(userSettings.Viewers), pq.Array(userSettings.Editors)).Scan(&id)
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		fmt.Printf("Inserted a single record %v\n", id)

	}
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

	// Execute(db, sqlStatement)
	Users := []models.User{
		{
			Name:           "lithial",
			Password:       "1234",
			Email:          "me@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "joe",
			Password:       "1234",
			Email:          "you@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "peter",
			Password:       "1234",
			Email:          "us@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "arran",
			Password:       "1234",
			Email:          "re@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "finn",
			Password:       "1234",
			Email:          "de@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "sam",
			Password:       "1234",
			Email:          "la@james.me",
			UserSettingsID: 0,
		},
	}

	for _, user := range Users {
		//fmt.Println(user.Name)
		var id int64
		canInsert := true
		u, _ := getUserByName(string(user.Name), db)
		if u.Name == user.Name {
			canInsert = false
			log.Printf("This user name is already taken\n")
		}
		if canInsert == true {
			sqlStatement := `INSERT INTO users (name, password,email,gender,token) VALUES ($1, $2, $3,$4,$5) RETURNING userId`
			//fmt.Printf("offending id = %d", )
			err := db.QueryRow(sqlStatement, user.Name, user.Password, user.Email, user.Gender, user.Token).Scan(&id)
			if err != nil {
				log.Printf("Unable to execute the query. %v\n", err)
			}
			fmt.Printf("Inserted a single record %v\n", id)
		}
	}
}
