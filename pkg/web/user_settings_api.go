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

//------------------------------JSON Webrequests User Settings Hander functions -- Users --------------------------------//

//ReturnSingleUserSettings Get usersetting in json format by id
//use mux to get us single notes
func (server Server) ReturnSingleUserSettings(w http.ResponseWriter, r *http.Request) {
	//we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we wish to return
	// convert the id type from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		res := models.BuildAPIResponseFail("Get Failed", nil)
		json.NewEncoder(w).Encode(res)

	} else {

		// call the getUser function with user id to retrieve a single user
		userSettings := getUserSettings(int64(id), server.db)
		if userSettings.ID == 0 {
			res := models.BuildAPIResponseFail("Get Failed", nil)
			json.NewEncoder(w).Encode(res)

		} else {
			res := models.BuildAPIResponseSuccess("Users retrieved successfully", userSettings)
			// send the response
			json.NewEncoder(w).Encode(res)
		}
	}

}

//CreateNewUserSettings Create UserSettings in json format
func (server Server) CreateNewUserSettings(w http.ResponseWriter, r *http.Request) {
	// return the string response containing the request body

	var userSettings models.UserSettings
	err := json.NewDecoder(r.Body).Decode(&userSettings)

	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
	}
	var res models.APIResponse
	// if user.viewers == "" || user.Email == "" || user.Password == "" {
	// 	res = models.BuildAPIResponseFail("Blank users cannot be created", nil)
	// } else {
	// call insert user function and pass the note
	insertID, _ := insertUserSettings(userSettings, server.db)
	// format a response object
	res = models.BuildAPIResponseSuccess(fmt.Sprintf("Usersettings Created with %d id", insertID), insertID)
	//}
	json.NewEncoder(w).Encode(res)

}

//DeleteUserSettings deletes a UserSettings
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
	// format the response message
	res := models.BuildAPIResponseSuccess(msg, deletedRows)
	// send the response
	json.NewEncoder(w).Encode(res)
}

//UpdateUserSettings updates the UpdateUserSettings as json
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
	err = json.NewDecoder(r.Body).Decode(&userSettings)

	userSettings.ID = id
	for i := 0; i < len(userSettings.Editors); i++ {
		if userSettings.Editors[i] == 0 {
			userSettings.Editors = append(userSettings.Editors[:i], userSettings.Editors[:i]...)
		}
	}
	for i := 0; i < len(userSettings.Viewers); i++ {
		if userSettings.Viewers[i] == 0 {
			userSettings.Viewers = append(userSettings.Viewers[:i], userSettings.Viewers[:i]...)
		}
	}
	// decode the json request to note
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		res := models.BuildAPIResponseFail("User Settings note saved", nil)
		json.NewEncoder(w).Encode(res)
	} else {
		// call update note to update the note
		updatedRows, _ := updateuserSettings(userSettings, server.db)
		log.Printf("Rows updated = %d", updatedRows)
		// format the message string
		//msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
		// format the response message
		userSettingsNew := getUserSettings(int64(userSettings.ID), server.db)
		res := models.BuildAPIResponseSuccess("User settings updated successfully", userSettingsNew)

		// send the response
		json.NewEncoder(w).Encode(res)
	}

}
