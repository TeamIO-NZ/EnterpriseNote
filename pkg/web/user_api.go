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

//ReturnSingleUserByEmail Get Notes in json format by username
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
	// format the response message
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
