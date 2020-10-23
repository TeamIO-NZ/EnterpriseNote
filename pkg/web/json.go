package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" //sql driver. blank is required
)

//------------------------------JSON Webrequests Hander functions--------------------------------//

//ReturnAllNotes Gets all the notes in json format
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

//ReturnSingleNote Get Notes with mux in json format
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

//CreateNewNote Create Note in json format
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
