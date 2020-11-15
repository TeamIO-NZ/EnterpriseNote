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
func (server Server) ReturnSingleUserSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 //we will need to parse the path parameters
	id, err := strconv.Atoi(vars["id"]) // convert the id type from string to int
	var res models.APIResponse
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		res = models.BuildAPIResponseFail("invalid id on User settings get", nil)
	} else {
		userSettings := getUserSettings(int64(id), server.db) // call the getUser function with user id to retrieve a single user
		if userSettings.ID == 0 {
			res = models.BuildAPIResponseFail("invalid User settings id", nil)
		} else {
			res = models.BuildAPIResponseSuccess("User settings retrieved successfully", userSettings)
		}
	}
	json.NewEncoder(w).Encode(res) // send the response
}

//CreateNewUserSettings Create UserSettings in json format
func (server Server) CreateNewUserSettings(w http.ResponseWriter, r *http.Request) {
	var userSettings models.UserSettings
	var res models.APIResponse
	err := json.NewDecoder(r.Body).Decode(&userSettings)
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		res = models.BuildAPIResponseFail("Unable to decode the request body", nil)
	} else {
		insertID, _ := insertUserSettings(userSettings, server.db)
		res = models.BuildAPIResponseSuccess(fmt.Sprintf("Usersettings Created with %d id", insertID), insertID)
	}
	json.NewEncoder(w).Encode(res)
}

//DeleteUserSettings deletes a UserSettings
func (server Server) DeleteUserSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 //mux params
	id, err := strconv.Atoi(vars["id"]) // convert the id in string to int
	var res models.APIResponse
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		res = models.BuildAPIResponseFail("Unable to convert the string into int.", nil)
	} else {
		deletedRows := deleteUserSettings(int64(id), server.db) // call the deleteUser, convert the int to int64
		res = models.BuildAPIResponseSuccess("User settings updated successfully.", deletedRows)
	}
	json.NewEncoder(w).Encode(res)
}

//UpdateUserSettings updates the UpdateUserSettings as json
func (server Server) UpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // mux params
	var res models.APIResponse
	id, err := strconv.Atoi(vars["id"]) // convert the id type from string to int
	if err != nil {
		log.Printf("Unable to convert the string into an int. %v", err)
		res = models.BuildAPIResponseFail("Invalid id passed to update user settings", nil)
	} else {
		var userSettings models.UserSettings                // create an empty user settings
		err = json.NewDecoder(r.Body).Decode(&userSettings) // decode the json request to note
		userSettings.ID = id                                // set the user settings id
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
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
			res = models.BuildAPIResponseFail("User Settings note saved", nil)
		} else {
			updateuserSettings(userSettings, server.db) // call update note to update the note
			userSettingsNew := getUserSettings(int64(userSettings.ID), server.db)
			res = models.BuildAPIResponseSuccess("User settings updated successfully", userSettingsNew)

		}
		json.NewEncoder(w).Encode(res) // send the response
	}

}
