package web

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"strconv"  // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route
	_ "github.com/lib/pq"    //sql driver. blank is required
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//------------------------------JSON Webrequests Hander functions -- Notes --------------------------------//

//ReturnAllNotes Gets all the notes in json format
func (server Server) ReturnAllNotes(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	//get all the notes in the database. returns the notes and any errors
	notes, err := getAllNotes()

	if err != nil {
		log.Fatalf("Unable to get all notes/ %v", err)
	}
	// send all the notes as response
	json.NewEncoder(w).Encode(notes)
}

//ReturnSingleNote Get Notes in json format by id
//use mux to get us single notes
func (server Server) ReturnSingleNote(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
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
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// return the string response containing the request body

	var note models.Note
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

//DeleteNote deletes a note
func (server Server) DeleteNote(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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

//UpdateNote updates the note as json
func (server Server) UpdateNote(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into an int. %v", err)
	}
	// create an empty note of type note
	var note models.Note

	// decode the json request to note
	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// call update note to update the note
	updatedRows := updateNote(int64(id), note)
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------------JSON Webrequests Hander functions -- Users --------------------------------//

//ReturnAllUsers Gets all the notes in json format
func (server Server) ReturnAllUsers(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	//get all the notes in the database. returns the notes and any errors
	users, err := getAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all notes/ %v", err)
	}
	// send all the notes as response
	json.NewEncoder(w).Encode(users)
}

//ReturnSingleUser Get Notes in json format by id
//use mux to get us single notes
func (server Server) ReturnSingleUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
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
	note, err := getUser(int64(id))

	if err != nil {
		log.Fatalf("Unable to get note. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(note)
}

//ReturnSingleUser Get Notes in json format by username
//use mux to get us single notes
func (server Server) ReturnSingleUserByName(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to return
	// convert the id type from string to int
	name := vars["username"]

	// call the getUser function with user id to retrieve a single user
	note, err := getUserByName(name)

	if err != nil {
		log.Fatalf("Unable to get note. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(note)
}

//CreateNewUser Create Note in json format
func (server Server) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// Allow all origin to handle cors issue
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// return the string response containing the request body

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// call insert user function and pass the note
	insertID := insertUser(user)
	// format a response object
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

//DeleteUser deletes a note
func (server Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we need to delete

	// convert the id in string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// call the deleteUser, convert the int to int64
	deletedRows := deleteUser(int64(id))

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

//UpdateUser updates the note as json
func (server Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into an int. %v", err)
	}
	// create an empty note of type note
	var user models.User

	// decode the json request to note
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// call update note to update the note
	updatedRows := updateUser(int64(id), user)
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------------JSON Webrequests Hander functions -- Specifics --------------------------------//

//SearchForNotes function that returns a bunch of notes with specific searching
func (server Server) SearchForNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into an int. %v", err)
	}
	//get specific notes in the database. returns the notes and any errors
	//input is 1-5 based on what notes we want. More functions to come maybe?
	notes, err := getSpecificNotes(id)

	if err != nil {
		log.Fatalf("Unable to get all notes/ %v", err)
	}
	// send all the notes as response
	json.NewEncoder(w).Encode(notes)

}

//GetAllNotesUserHasAccessTo function that returns a bunch of notes with specific searching
func (server Server) GetAllNotesUserHasAccessTo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into an int. %v", err)
	}
	//var notes []models.Note
	//get specific notes in the database. returns the notes and any errors
	//input is 1-5 based on what notes we want. More functions to come maybe?
	notes, err := getAllNotesUserHasAccessTo(id)

	if err != nil {
		log.Fatalf("Unable to get all notes/ %v", err)
	}
	// send all the notes as response
	json.NewEncoder(w).Encode(notes)

}

//Login method that generates an api key and returns it to the client if the provided login information is correct
func (server Server) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	username := r.Header.Get("X-Username")
	password := r.Header.Get("X-Password")
	if username == "" || password == "" {
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Code:    400,
			Message: "400: Bad Request. Empty username or password, set X-Username or X-Password headers to use.",
		})
	}
	//TODO: check if creds are valid
	//TODO: generate api token
	//TODO: return api token
	//create the data
	var data, err = checkLogin(username, password)
	//check for errors
	if err != nil {
		log.Printf("Unable to get user login/ %v", err)
	}
	//encode the data
	json.NewEncoder(w).Encode(data)
	// _ = json.NewEncoder(w).Encode(models.APIResponse{
	// 	Code: 200,
	// 	Data: "totally-an-api-key",
	// })
}
