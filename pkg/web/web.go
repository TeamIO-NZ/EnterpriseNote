package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

//Start this is the function that starts the webserver
func (server Server) Start() {
	zap.S().Info("Starting webserver...")
	server.config.Port = "8080"
	if server.config.Port == "" {
		server.config.Port = "8080"
		//zap.S().Warn("No webserver port config detected, using 8080.")
	}
	Notes = []Note{
		Note{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Note{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	server.handleRequests()

}
func (server Server) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homepage")
}

//------------------------------JSON Webrequests Hander functions--------------------------------//

//Get All Note
func (server Server) ReturnAllNotes(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("EndpointHit: return all notes")

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	notes, err := getAllNotes()

	if err != nil {
		log.Fatalf("Unable to get all notes/ %v", err)
	}
	json.NewEncoder(w).Encode(notes)

}

//Get Note
//use mux to get us single notes
func (server Server) ReturnSingleNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to return
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	//key := vars["id"]
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// call the getUser function with user id to retrieve a single user
	note, err := getNote(int64(id))

	if err != nil {
		log.Fatalf("Unable to get note. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(note)
}

//Create Note
func (server Server) CreateNewNote(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return the string response containing the request body

	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// call insert user function and pass the note
	insertID := insertNote(note)
	// format a response object
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

//Delete Note
func (server Server) DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we need to delete

	// convert the id in string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// call the deleteUser, convert the int to int64
	deletedRows := deleteNote(int64(id))

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)
	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	// send the response
	json.NewEncoder(w).Encode(res)
}

//Update Note
func (server Server) UpdateNote(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into an int. %v", err)
	}
	var note Note

	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateNote(int64(id), note)

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------------Main Request Handler functions--------------------------------//

//run me to make the server work
func (server Server) handleRequests() {

	createTable()
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/notes", server.ReturnAllNotes).Methods("GET")
	r.HandleFunc("/api/v1/note/{id}", server.ReturnSingleNote).Methods("GET")
	r.HandleFunc("/api/v1/note", server.CreateNewNote).Methods("POST")
	r.HandleFunc("/api/v1/note/{id}", server.UpdateNote).Methods("PUT")

	r.HandleFunc("/api/v1/note/{id}", server.DeleteNote).Methods("DELETE")
	r.Handle("/", http.RedirectHandler("/web/", http.StatusPermanentRedirect))
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	log.Fatal(http.ListenAndServe(":8080", r))
}

//getNote
func getNote(id int64) (Note, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var note Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&note.ID, &note.Title, &note.Desc, &note.Content)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return note, nil
	case nil:
		return note, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return note, err
}

// get one user from the DB by its userid
func getAllNotes() ([]Note, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var notes []Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var note Note

		// unmarshal the row object to user
		err = rows.Scan(&note.ID, &note.Title, &note.Desc, &note.Content)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		notes = append(notes, note)

	}

	// return empty user on error
	return notes, err
}

// update user in the DB
func updateNote(id int64, note Note) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE notes SET title=$2, description=$3, contents=$4 WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, note.Title, note.Desc, note.Content)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func deleteNote(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM notes WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

//insert note
func insertNote(note Note) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO notes (title, description, contents) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, note.Title, note.Desc, note.Title).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

//------------------------------SQL Hander functions--------------------------------//

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
