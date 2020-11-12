package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

//User struct declaration
type User struct {
	gorm.Model
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"type:varchar(100);unique_index"`
	Gender         string `json:"Gender"`
	Password       string `json:"Password"`
	Token          string `json:"Token"`
	UserSettingsID int    `json:UserSettingsId`
}

//User common settings table
type UserSettings struct {
	ID      int   `Json:"id`
	Viewers []int `json:"Viewers"`
	Editors []int `json:"Editors"`
}

//ParseSingleUser does a thing
func ParseSingleUser(row *sql.Rows) (user User, err error) {

	//fmt.Println("scanning a row of user stuff. awaiting crash")
	if row.Next() {
		// unmarshal the row object to user
		err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Gender, &user.Email, &user.Token)
		if err != nil {
			log.Printf("ParseSingleNote: Unable to scan the row. %v\n", err)
			return User{}, err
		}
	}
	return user, nil
}

//ParseSingleNote Parses a single note and returns it
func ParseSingleUserSetting(row *sql.Rows) UserSettings {
	var userSettings UserSettings
	var viewers string
	var editors string
	fmt.Println("scanning a row of note stuff. awaiting crash")
	// unmarshal the row object to user
	if row.Next() {
		err := row.Scan(&userSettings.ID, &viewers, &editors)
		if err != nil {
			log.Printf("ParseSingleNote: Unable to scan the row. %v\n", err)
		}
		userSettings.Viewers = ParseStringForArrayNumbers(viewers)
		userSettings.Editors = ParseStringForArrayNumbers(editors)
	}
	//return the note
	return userSettings
}

//ParseUserArray parses a user array and returns the entire array
func ParseUserArray(rows *sql.Rows) []User {
	var user User
	var users []User
	for rows.Next() {
		//fmt.Println("scanning a row of user stuff. awaiting crash")
		// unmarshal the row object to user
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Gender, &user.Email, &user.Token, &user.UserSettingsID)
		if err != nil {
			log.Printf("ParseUserArray: Unable to scan the row. %v\n", err)
		}
		// // append the user in the users slice
		users = append(users, user)
	}
	return users
}
