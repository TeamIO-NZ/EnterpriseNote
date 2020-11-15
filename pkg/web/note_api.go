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
	//get all the notes in the database. returns the notes in a response
	var res models.APIResponse
	notes := getAllNotes(server.db)
	if len(notes) == 0 {
		res = models.BuildAPIResponseFail("No notes found in this database", nil)
	} else {
		res = models.BuildAPIResponseSuccess("Notes found", notes)
	}
	json.NewEncoder(w).Encode(res)
}

//ReturnSingleNote Get Notes in json format by id
func (server Server) ReturnSingleNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 //parse the params
	id, err := strconv.Atoi(vars["id"]) //extra the id and make it a string
	if err != nil {                     //error check for the conversion
		log.Printf("Unable to convert the string into int.  %v", err)
	}
	note := getNote(int64(id), server.db) // get the note
	fmt.Println(note)
	var res models.APIResponse
	if note.Content == "" {
		res = models.BuildAPIResponseFail("Note with that not found", nil)
	} else {
		res = models.BuildAPIResponseSuccess("Note found", note)
	}
	json.NewEncoder(w).Encode(res) // send the response
}

//CreateNewNote Create Note from incoming json
func (server Server) CreateNewNote(w http.ResponseWriter, r *http.Request) {

	// return the string response containing the request body
	var note models.Note
	var res models.APIResponse

	err := json.NewDecoder(r.Body).Decode(&note) //decode sent data
	if err != nil {                              //error check
		log.Printf("Unable to decode the request body.  %v", err)
		res = models.BuildAPIResponseFail("Invalid Json when creating new note", nil) // fail response
	} else {
		insertID := insertNote(note, server.db)                                     // call insert user function and pass the note
		res = models.BuildAPIResponseSuccess("Note created successfully", insertID) // success response
	}
	json.NewEncoder(w).Encode(res) //send response

}

//DeleteNote deletes a note with the given id
func (server Server) DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 //mux params
	id, err := strconv.Atoi(vars["id"]) // convert the id in string to int
	var res models.APIResponse

	if err != nil { //error checking
		log.Printf("Unable to convert the string into int.  %v", err)
		res = models.BuildAPIResponseFail("Invalid id, note doesn't exist", nil)
	} else {
		deletedRows := deleteNote(id, server.db) // call the deleteUser
		res = models.BuildAPIResponseSuccess("User deleted successfully", deletedRows)
	}
	json.NewEncoder(w).Encode(res) // send the response

}

//UpdateNote updates the note as with the given id using the given json
func (server Server) UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 // get the userid from the request params, key is "id"
	id, err := strconv.Atoi(vars["id"]) // convert the id type from string to int
	var res []models.APIResponse

	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
		res = append(res, models.BuildAPIResponseFail("param id is note an int", nil))
	}
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
		res = append(res, models.BuildAPIResponseFail("param userid is note an int", nil))
	}
	var canChange = false                       //used to determine if the change request comes from a valid user
	var note models.Note                        // create an empty note of type note
	err = json.NewDecoder(r.Body).Decode(&note) // decode the json request to note
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		res = append(res, models.BuildAPIResponseFail("Invalid json when updating a note", nil))
	}
	if note.Owner == userid {
		canChange = true
	} else {
		for _, a := range note.Editors {
			if a == userid {
				canChange = true
				break
			}
		}
	}
	if !canChange {
		res = append(res, models.BuildAPIResponseUnauthorized("You are not authorized to do this"))
	} else {
		res = res[:0]
		updatedRows := updateNote(int64(id), note, server.db)                                       // call update note to update the note
		res = append(res, models.BuildAPIResponseSuccess("User updated successfully", updatedRows)) // format the response message
	}
	json.NewEncoder(w).Encode(res) // send the response
}

//GetAllNotesUserHasAccessTo function that returns a bunch of notes the specific user has access to
func (server Server) GetAllNotesUserHasAccessTo(w http.ResponseWriter, r *http.Request) {
	var res models.APIResponse
	vars := mux.Vars(r)                 //params
	id, err := strconv.Atoi(vars["id"]) // convert the id type from string to int
	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
		res = models.BuildAPIResponseFail("Invalid id, note doesn't exist", nil)
	} else {
		notes := getAllNotesUserHasAccessTo(id, server.db)
		res = models.BuildAPIResponseSuccess("Some notes were found", notes)
	}
	json.NewEncoder(w).Encode(res) // send all the notes as response

}
