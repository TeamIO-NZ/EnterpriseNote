package models

import (
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
)

//User struct declaration
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"type:varchar(100);unique_index"`
	Gender   string `json:"Gender"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
}

//ParseSingleUser does a thing
func ParseSingleUser(row *sql.Rows) User {
	var user User

	//fmt.Println("scanning a row of user stuff. awaiting crash")
	if row.Next() {
		// unmarshal the row object to user
		err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Token)
		if err != nil {
			log.Printf("ParseSingleNote: Unable to scan the row. %v\n", err)
		}
	}
	return user
}

//ParseUserArray parses a user array and returns the entire array
func ParseUserArray(rows *sql.Rows) []User {
	var user User
	var users []User
	for rows.Next() {
		//fmt.Println("scanning a row of user stuff. awaiting crash")
		// unmarshal the row object to user
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Token)
		if err != nil {
			log.Printf("ParseUserArray: Unable to scan the row. %v\n", err)
		}
		// // append the user in the users slice
		users = append(users, user)
	}
	return users
}
