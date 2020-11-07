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
	// Notes = []models.Note{
	// 	models.Note{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	// 	models.Note{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	// }
	server.db = createConnection()
	server.HandleRequests()
	server.db.Close()
}

//HandleRequests run me to make the server work
func (server Server) HandleRequests() {

	createTable()
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/notes", server.ReturnAllNotes).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.ReturnSingleNote).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/note", server.CreateNewNote).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.UpdateNote).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/note/{id}", server.DeleteNote).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/users", server.ReturnAllUsers).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user/{username}", server.ReturnSingleUserByName).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.ReturnSingleUser).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/user", server.CreateNewUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.UpdateUser).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/user/{id}", server.DeleteUser).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/usersnotes/{id}", server.GetAllNotesUserHasAccessTo).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/login", server.Login).Methods("GET")

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

//TODO function for dropping and rebuilding tables.
//TODO function for loading in a csv for a temp table or something
func createTable() {
	//creates database connection
	db := createConnection()
	//prepares to close database when done
	defer db.Close()

	//create the base notes table for if it doesn't exist
	sqlStatement := `DROP TABLE IF EXISTS notes;`
	Execute(db, sqlStatement)
	// //execute the sql statement and return a response
	// res, err := db.Exec(sqlStatement)
	// if err != nil {
	// 	log.Fatalf("Unable to execute the query. %v", err)
	// }
	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP TABLE IF EXISTS users;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	// res, err; = db.Exec(sqlStatement)
	// if err != nil {
	// 	log.Fatalf("Unable to execute the query. %v", err)
	// }

	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		password TEXT,
		email TEXT,
		token TEXT
	);`
	Execute(db, sqlStatement)

	//execute the sql statement and return a response
	// res, err = db.Exec(sqlStatement)
	// if err != nil {
	// 	log.Fatalf("Unable to execute the query. %v", err)
	// }
	//print the response maybe
	//fmt.Printf("%s\n ", res)

	//TODO rebuild this table function
	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS notes (
			id SERIAL PRIMARY KEY,
			title TEXT,
			description TEXT,
			contents TEXT,
			owner INT,
			viewers integer[],
			editors integer[],
			FOREIGN KEY (owner)	REFERENCES users (id)
		);`
	Execute(db, sqlStatement)

	//execute the sql statement and return a response
	// res, err = db.Exec(sqlStatement)
	// if err != nil {
	// 	log.Fatalf("Unable to execute the query. %v", err)
	// }

	// sqlStatement = `
	// insert into users (id,name,password,email) values (0,'Lithial','1234','me@james.me');
	// insert into users (id,name,password,email) values (1,'Joe','1234','you@james.me');
	// insert into users (id,name,password,email) values (2,'Peter','1234','us@james.me');
	// insert into users (id,name,password,email) values (3,'Arran','1234','re@james.me');
	// insert into users (id,name,password,email) values (4,'Finn','1234','la@james.me');
	// insert into users (id,name,password,email) values (5,'Sam','1234','de@james.me');
	// insert into users (id,name,password,email) values (6,'Sam','1234','de@james.me');
	// `
	// Execute(db, sqlStatement)
	Users := []models.User{
		models.User{
			Name:     "Lithial",
			Password: "1234",
			Email:    "me@james.me",
		},
		models.User{
			Name:     "Joe",
			Password: "1234",
			Email:    "you@james.me",
		},
		models.User{
			Name:     "Peter",
			Password: "1234",
			Email:    "us@james.me",
		},
		models.User{
			Name:     "Arran",
			Password: "1234",
			Email:    "re@james.me",
		},
		models.User{
			Name:     "Finn",
			Password: "1234",
			Email:    "de@james.me",
		},
		models.User{
			Name:     "Sam",
			Password: "1234",
			Email:    "la@james.me",
		},
		models.User{
			Name:     "Sam",
			Password: "1234",
			Email:    "ke@james.me",
		},
	}

	for _, user := range Users {
		//fmt.Println(user.Name)
		var id int64
		canInsert := true
		u := getUserByName(string(user.Name), db)
		if u.Name == user.Name {
			canInsert = false
			log.Printf("This user name is already taken\n")
		}
		if canInsert == true {
			sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
			err := db.QueryRow(sqlStatement, user.Name, user.Password, user.Email).Scan(&id)
			//TODO make this error message less bad
			if err != nil {
				log.Printf("Unable to execute the query. %v\n", err)
			}
			fmt.Printf("Inserted a single record %v\n", id)
		}
	}

	Notes = []models.Note{
		models.Note{
			ID:      "0",
			Title:   "James is the overlord",
			Desc:    "The best overlord",
			Content: "The very best overlord there is",
			Owner:   0,
			Viewers: []int{1, 2, 3},
			Editors: []int{4, 5},
		},
		models.Note{
			ID:      "1",
			Title:   "Joe is the Minion",
			Desc:    "The best minion",
			Content: "So i decree",
			Owner:   0,
			Viewers: []int{1, 2, 3},
			Editors: []int{4, 5},
		},
		models.Note{
			ID:      "2",
			Title:   "No joe is the boss",
			Desc:    "The best boss",
			Content: "So i decree",
			Owner:   1,
			Viewers: []int{0, 2, 3},
			Editors: []int{4, 5},
		},
	}
	for _, note := range Notes {
		fmt.Println(note.Desc)
		var id int64
		sqlStatement := `INSERT INTO notes (id, title, description, contents, owner, viewers, editors) VALUES ($1, $2, $3,$4,$5,$6, $7) RETURNING id`

		err := db.QueryRow(sqlStatement, note.ID, note.Title, note.Desc, note.Content, note.Owner, pq.Array(note.Viewers), pq.Array(note.Editors)).Scan(&id)
		//TODO make this error message less bad
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		fmt.Printf("Inserted a single record %v\n", id)

	}
}
