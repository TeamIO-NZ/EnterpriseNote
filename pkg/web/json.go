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
	//get all the notes in the database. returns the notes and any errors
	notes := getAllNotes(server.db)

	// send all the notes as response
	json.NewEncoder(w).Encode(notes)
}

//ReturnSingleNote Get Notes in json format by id
//use mux to get us single notes
func (server Server) ReturnSingleNote(w http.ResponseWriter, r *http.Request) {
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we wish to return
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
	}
	// call the getUser function with user id to retrieve a single user
	note := getNote(int64(id), server.db)
	fmt.Println(note)
	// send the response
	json.NewEncoder(w).Encode(note)
}

//CreateNewNote Create Note in json format
func (server Server) CreateNewNote(w http.ResponseWriter, r *http.Request) {

	// return the string response containing the request body
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
	}
	fmt.Println(note)
	// call insert user function and pass the note
	insertID, err := insertNote(note, server.db)
	// format a response object
	// res := response{
	// 	ID:      insertID,
	// 	Message: "Note created successfully",
	// }
	res := models.BuildAPIResponseSuccess("User created successfully", insertID)
	json.NewEncoder(w).Encode(res)

}

//DeleteNote deletes a note
func (server Server) DeleteNote(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting note start")
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we need to delete
	// convert the id in string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
	}
	log.Println("Deleting note")
	// call the deleteUser, convert the int to int64
	deletedRows := deleteNote(id, server.db)

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

	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
	}
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
	}
	fmt.Println("This is an updating of the note")
	var canChange = false
	// create an empty note of type note
	var note models.Note
	// decode the json request to note
	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
	}
	if note.Owner == userid {
		fmt.Println("I am the owner")
		canChange = true
	} else {
		fmt.Println("am i an editor")

		for _, a := range note.Editors {
			if a == userid {
				fmt.Println("oh i am")
				canChange = true
				break
			}
		}
	}
	if canChange {
		// call update note to update the note
		updatedRows := updateNote(int64(id), note, server.db)
		// format the message string
		msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
		// format the response message
		res := response{
			ID:      int64(id),
			Message: msg,
		}
		// send the response
		json.NewEncoder(w).Encode(res)
	} else {
		res := models.BuildAPIResponseUnauthorized("You are not authorized to do this")
		// send the response
		json.NewEncoder(w).Encode(res)
	}

}

//------------------------------JSON Webrequests Hander functions -- Users --------------------------------//

//ReturnAllUsers Gets all the notes in json format
func (server Server) ReturnAllUsers(w http.ResponseWriter, r *http.Request) {
	//get all the notes in the database. returns the notes and any errors
	users := getAllUsers(server.db)
	// send all the notes as response
	json.NewEncoder(w).Encode(users)
}

//ReturnSingleUser Get Notes in json format by id
//use mux to get us single notes
func (server Server) ReturnSingleUser(w http.ResponseWriter, r *http.Request) {
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we wish to return
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
	}
	// call the getUser function with user id to retrieve a single user
	note := getUser(int64(id), server.db)

	// send the response
	json.NewEncoder(w).Encode(note)
}

//ReturnSingleUserByName Get Notes in json format by username
//use mux to get us single notes
func (server Server) ReturnSingleUserByName(w http.ResponseWriter, r *http.Request) {
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we wish to return
	// convert the id type from string to int
	name := vars["username"]

	// call the getUser function with user id to retrieve a single user
	note, err := getUserByName(name, server.db)
	if err != nil {
		json.NewEncoder(w).Encode(models.BuildAPIResponseFail("Username in use", nil))
		return
	}

	// send the response
	json.NewEncoder(w).Encode(note)
}

//ReturnSingleUserByName Get Notes in json format by username
//use mux to get us single notes
func (server Server) ReturnSingleUserByEmail(w http.ResponseWriter, r *http.Request) {
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we wish to return
	// convert the id type from string to int
	email := vars["email"]

	// call the getUser function with user id to retrieve a single user
	user, err := getUserByEmail(email, server.db)
	if err != nil {
		json.NewEncoder(w).Encode(models.BuildAPIResponseFail("Email in use", nil))
		return
	}

	// send the response
	json.NewEncoder(w).Encode(user)

}

//CreateNewUser Create Note in json format
func (server Server) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	// return the string response containing the request body

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
	}
	var res models.APIResponse
	if user.Name == "" || user.Email == "" || user.Password == "" {
		res = models.BuildAPIResponseFail("Blank users cannot be created", nil)
	} else {
		// call insert user function and pass the note
		insertID := testInsert(user, server.db)
		// format a response object
		res = models.BuildAPIResponseSuccess(fmt.Sprintf("User Created with %d id", insertID), nil)
	}
	json.NewEncoder(w).Encode(res)

}

//DeleteUser deletes a note
func (server Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we need to delete

	// convert the id in string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
	}
	// call the deleteUser, convert the int to int64
	deletedRows := deleteUser(int64(id), server.db)

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
	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
	}
	// create an empty note of type note
	var user models.User

	// decode the json request to note
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
	}
	// call update note to update the note
	updatedRows := updateUser(int64(id), user, server.db)
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

//------------------------------JSON Webrequests User Settings Hander functions -- Users --------------------------------//

//ReturnSingleUser Get Notes in json format by id
//use mux to get us single notes
func (server Server) ReturnSingleUserSettings(w http.ResponseWriter, r *http.Request) {
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we wish to return
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
	}
	// call the getUser function with user id to retrieve a single user
	note := getUserSettings(int64(id), server.db)

	// send the response
	json.NewEncoder(w).Encode(note)
}

//CreateNewUser Create Note in json format
func (server Server) CreateNewUserSettings(w http.ResponseWriter, r *http.Request) {
	// return the string response containing the request body

	var userSettings models.UserSettings
	err := json.NewDecoder(r.Body).Decode(&userSettings)
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		res := models.BuildAPIResponseFail("User Settings note saved", nil)
		json.NewEncoder(w).Encode(res)
	} else {
		insertID := insertUserSettings(userSettings, server.db)
		// format a response object
		res := models.BuildAPIResponseSuccess("User settings created succesfully", insertID)
		json.NewEncoder(w).Encode(res)
	}
	// call insert user function and pass the note

}

//DeleteUser deletes a note
func (server Server) DeleteUserSettings(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we need to delete

	// convert the id in string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
	}
	// call the deleteUser, convert the int to int64
	deletedRows := deleteUserSettings(int64(id), server.db)

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
func (server Server) UpdateUserSettings(w http.ResponseWriter, r *http.Request) {

	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
	}
	// create an empty note of type note
	var userSettings models.UserSettings

	// decode the json request to note
	err = json.NewDecoder(r.Body).Decode(&userSettings)
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
	}
	// call update note to update the note
	updatedRows := updateuserSettings(int64(id), userSettings, server.db)
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

//GetAllNotesUserHasAccessTo function that returns a bunch of notes with specific searching
func (server Server) GetAllNotesUserHasAccessTo(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
	}
	//get specific notes in the database. returns the notes and any errors
	//input is 1-5 based on what notes we want. More functions to come maybe?
	notes := getAllNotesUserHasAccessTo(id, server.db)

	// send all the notes as response
	json.NewEncoder(w).Encode(notes)

}

//SearchNotesForSpecifics //TODO this needs work. it currently returns no rows
func (server Server) SearchNotesForSpecifics(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request params, key is "id"
	vars := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	prefix := vars["prefix"]
	targetFunction := vars["targetFunction"]

	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
	}
	//get specific notes in the database. returns the notes and any errors
	//input is 1-5 based on what notes we want. More functions to come maybe?
	notes := getSpecificNotes(server.db, id, targetFunction, prefix)
	// send all the notes as response
	json.NewEncoder(w).Encode(notes)
}

//Login method that generates an api key and returns it to the client if the provided login information is correct
func (server Server) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]
	if username == "" || password == "" {
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Code:    400,
			Message: "400: Bad Request. Empty username or password, set X-Username or X-Password headers to use.",
		})
	}
	//create the data
	var data = checkLogin(username, password, server.db)

	//encode the data
	json.NewEncoder(w).Encode(data)
}

func (server Server) CheckConnection(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.BuildAPIResponseSuccess("Success", nil))
}
