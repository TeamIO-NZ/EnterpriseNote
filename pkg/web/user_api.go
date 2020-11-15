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
	users := getAllUsers(server.db) //get all the notes in the database. returns the notes and any errors
	var res models.APIResponse
	if len(users) == 0 {
		res = models.BuildAPIResponseFail("Found no users", nil)
	} else {
		res = models.BuildAPIResponseSuccess("Found all users", users)
	}
	json.NewEncoder(w).Encode(res) // send all the notes as response
}

//ReturnSingleUser Get Notes in json format by id
func (server Server) ReturnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 // mux params
	id, err := strconv.Atoi(vars["id"]) // convert the id type from string to int
	var res models.APIResponse

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		res = models.BuildAPIResponseFail("Invalid id", nil)
	} else {
		note := getUser(int64(id), server.db) // call the getUser function with user id to retrieve a single user
		res = models.BuildAPIResponseSuccess("User is found", note)
	}
	json.NewEncoder(w).Encode(res) // send the response
}

//ReturnSingleUserByName Get Notes in json format by username
func (server Server) ReturnSingleUserByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)      //mux params
	name := vars["username"] //get the id
	var res models.APIResponse

	note, err := getUserByName(name, server.db) // call the getUser function with user id to retrieve a single user
	if err != nil {
		res = models.BuildAPIResponseFail("Username in use", nil)
	} else {
		res = models.BuildAPIResponseSuccess("User found", note)
	}
	json.NewEncoder(w).Encode(res) // send the response
}

//ReturnSingleUserByEmail Get Notes in json format by username
func (server Server) ReturnSingleUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                           // mux params
	email := vars["email"]                        // get the email value
	user, err := getUserByEmail(email, server.db) // call the getUser function with user id to retrieve a single user
	var res models.APIResponse
	if err != nil {
		res = models.BuildAPIResponseFail("Email in use", nil)
	} else {
		res = models.BuildAPIResponseSuccess("User found", user)
	}
	json.NewEncoder(w).Encode(res) // send the response
}

//CreateNewUser Create Note in json format
func (server Server) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User       // make a user
	var res models.APIResponse // make a response

	err := json.NewDecoder(r.Body).Decode(&user) //decode the user
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		res = models.BuildAPIResponseFail("Unable to decode the request body", nil)
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		res = models.BuildAPIResponseFail("Blank users cannot be created", nil)
	} else {
		insertID := insertUser(user, server.db)                                                     // call insert user function and pass the note
		res = models.BuildAPIResponseSuccess(fmt.Sprintf("User Created with %d id", insertID), nil) // format a response object
	}
	json.NewEncoder(w).Encode(res)

}

//DeleteUser deletes a note
func (server Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 //mux params
	id, err := strconv.Atoi(vars["id"]) // convert the id in string to int
	var res models.APIResponse          // make a response

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		res = models.BuildAPIResponseFail("Unable to convert the string into int", nil)
	} else {
		deletedRows := deleteUser(int64(id), server.db) // call the deleteUser, convert the int to int64
		res = models.BuildAPIResponseSuccess("User updated successfully.", deletedRows)
	}
	// send the response
	json.NewEncoder(w).Encode(res)
}

//UpdateUser updates the note as json
func (server Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 // mux params
	id, err := strconv.Atoi(vars["id"]) // convert the id type from string to int
	var res models.APIResponse          // make a response
	var user models.User                // make a user

	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
		res = models.BuildAPIResponseFail("Unable to convert the string into an int.", nil)
	} else {
		err = json.NewDecoder(r.Body).Decode(&user) // decode the json request to note
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
			res = models.BuildAPIResponseFail("Unable to decode the request body.", nil)
		} else {
			updatedRows := updateUser(int64(id), user, server.db) // call update note to update the note
			res = models.BuildAPIResponseSuccess("User updated successfully.", updatedRows)
		}
	}
	json.NewEncoder(w).Encode(res) // send the response
}

//Login method that generates an api key and returns it to the client if the provided login information is correct
func (server Server) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)          // mux params
	username := vars["username"] // get username
	password := vars["password"] // get password
	var res models.APIResponse   // make a response

	if username == "" || password == "" {
		res = models.BuildAPIResponseFail("Bad Request. Empty username or password", nil)
	} else {
		var data = checkLogin(username, password, server.db) //create the data
		res = models.BuildAPIResponseSuccess("Login successful", data)
	}
	json.NewEncoder(w).Encode(res) //encode the data

}
