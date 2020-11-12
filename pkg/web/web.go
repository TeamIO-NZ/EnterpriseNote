package web

import (
	"database/sql"
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable

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

	server.db = createConnection()
	server.HandleRequests()
	server.db.Close()
}

//HandleRequests run me to make the server work
func (server Server) HandleRequests() {

	//createTable()
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/notes", server.ReturnAllNotes).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.ReturnSingleNote).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/note", server.CreateNewNote).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}/{userid}", server.UpdateNote).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.DeleteNote).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}/{targetFunction}/{prefix}", server.SearchNotesForSpecifics).Methods("GET", "OPTIONS")

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
	db, err := sql.Open("postgres", os.Getenv("ALT_POSTGRES_URL"))
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

//populates database with fake data
func createTable() {
	//creates database connection
	db := createConnection()
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
	sqlStatement = `DROP SEQUENCE IF EXISTS user_sequence;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS notes_sequence;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)
	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE SEQUENCE user_sequence
	minvalue 0
	start 0
	increment 1;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE SEQUENCE notes_sequence
	minvalue 0
	start 0
	increment 1;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)
	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP TABLE IF EXISTS userSettings;`
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS userSettings (
			id int PRIMARY KEY,
			viewers integer[],
			editors integer[]
		);`
	Execute(db, sqlStatement)
	UserSettings := []models.UserSettings{
		models.UserSettings{
			ID:      0,
			Viewers: []int{0, 2, 3},
			Editors: []int{4, 5},
		},
		models.UserSettings{
			ID:      1,
			Viewers: []int{0, 2, 3},
			Editors: []int{4, 5},
		},
		models.UserSettings{
			ID:      2,
			Viewers: []int{0, 2, 3},
			Editors: []int{4, 5},
		},
	}
	for _, userSettings := range UserSettings {
		fmt.Println(userSettings.ID)
		var id int64
		sqlStatement := `INSERT INTO userSettings (id, viewers,editors) VALUES ($1,$2, $3) RETURNING id`

		err := db.QueryRow(sqlStatement, userSettings.ID, pq.Array(userSettings.Viewers), pq.Array(userSettings.Editors)).Scan(&id)
		//TODO make this error message less bad
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		fmt.Printf("Inserted a single record %v\n", id)

	}
	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS users (
		id int PRIMARY KEY,
		name TEXT,
		password TEXT,
		gender TEXT,
		email TEXT,
		token TEXT,
		userSettingsId int,
		FOREIGN KEY (userSettingsId) REFERENCES userSettings (id) on delete cascade on update cascade

	);`
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS notes (
			id int PRIMARY KEY,
			title TEXT,
			description TEXT,
			contents TEXT,
			owner INT,
			viewers integer[],
			editors integer[],
			FOREIGN KEY (owner)	REFERENCES users (id) on delete cascade on update cascade
		);`
	Execute(db, sqlStatement)

	// Execute(db, sqlStatement)
	Users := []models.User{
		models.User{
			Name:     "lithial",
			Password: "1234",
			Email:    "me@james.me",
		},
		models.User{
			Name:     "joe",
			Password: "1234",
			Email:    "you@james.me",
		},
		models.User{
			Name:     "peter",
			Password: "1234",
			Email:    "us@james.me",
		},
		models.User{
			Name:     "arran",
			Password: "1234",
			Email:    "re@james.me",
		},
		models.User{
			Name:     "finn",
			Password: "1234",
			Email:    "de@james.me",
		},
		models.User{
			Name:     "sam",
			Password: "1234",
			Email:    "la@james.me",
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
			sqlStatement := `INSERT INTO users (id, name, password,email,gender,token) VALUES (nextval('user_sequence'),$1, $2, $3,$4,$5) RETURNING id`
			//fmt.Printf("offending id = %d", )
			err := db.QueryRow(sqlStatement, user.Name, user.Password, user.Email, user.Gender, user.Token).Scan(&id)
			//TODO make this error message less bad
			if err != nil {
				log.Printf("Unable to execute the query. %v\n", err)
			}
			fmt.Printf("Inserted a single record %v\n", id)
		}
	}

	Notes = []models.Note{
		models.Note{
			Title:   "James is the overlord",
			Desc:    "The best overlord",
			Content: "The very best overlord there is",
			Owner:   0,
			Viewers: []int{1, 2, 3},
			Editors: []int{4, 5},
		},
		models.Note{
			Title:   "Joe is the Minion",
			Desc:    "The best minion",
			Content: "So i decree",
			Owner:   1,
			Viewers: []int{1, 2, 3},
			Editors: []int{4, 5},
		},
		models.Note{
			Title:   "No joe is the boss",
			Desc:    "The best boss",
			Content: "So i decree",
			Owner:   2,
			Viewers: []int{0, 2, 3},
			Editors: []int{4, 5},
		},
	}
	for _, note := range Notes {
		fmt.Println(note.Desc)
		var id int64
		sqlStatement := `INSERT INTO notes (id, title, description, contents, owner, viewers, editors) VALUES (nextval('notes_sequence'),$1,$2, $3,$4,$5,$6) RETURNING id`

		err := db.QueryRow(sqlStatement, note.Title, note.Desc, note.Content, note.Owner, pq.Array(note.Viewers), pq.Array(note.Editors)).Scan(&id)
		//TODO make this error message less bad
		if err != nil {
			log.Printf("note: %s is the offending note", note.Title)
			log.Printf("Unable to execute the query. %v\n", err)
		}
		fmt.Printf("Inserted a single record %v\n", id)

	}

}
