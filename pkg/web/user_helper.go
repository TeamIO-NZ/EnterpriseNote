package web

import (
	"database/sql"
	"log"

	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//------------------------------Main Request Handler functions- Users-------------------------------//

// get one user from the DB by its userid
func getUser(id int64, db *sql.DB) models.User {
	PingOrPanic(db)                                       //check connection
	sqlStatement := `SELECT * FROM users WHERE userId=$1` // create the select sql query
	row := QueryRowForType(db, sqlStatement, id)          // execute the sql statement
	user, _ := models.ParseSingleUser(row)                // unmarshal the row object to user
	defer row.Close()                                     //remember to close the rows
	return user                                           // return empty user on error
}

// get one user from the DB by its userid
func getUserByName(name string, db *sql.DB) (models.User, error) {
	PingOrPanic(db)                                     // check the connection
	sqlStatement := `SELECT * FROM users WHERE name=$1` // create the select sql query
	row := QueryRowForType(db, sqlStatement, name)      // execute the sql statement
	user, err := models.ParseSingleUser(row)            // unmarshal the row object to user
	defer row.Close()                                   //remember to close the rows
	return user, err                                    // return empty user on error
}

// get one user from the DB by its userid
func getUserByEmail(email string, db *sql.DB) (models.User, error) {
	PingOrPanic(db)                                      // check the connection
	sqlStatement := `SELECT * FROM users WHERE email=$1` // create the select sql query
	row := QueryRowForType(db, sqlStatement, email)      // execute the sql statement
	user, err := models.ParseSingleUser(row)             // unmarshal the row object to user
	defer row.Close()                                    //remember to close the rows
	return user, err                                     // return empty user on error
}

// get one user from the DB by its userid
func getAllUsers(db *sql.DB) []models.User {
	PingOrPanic(db)                           // check the connection
	var users []models.User                   // create user array
	sqlStatement := `SELECT * FROM users`     // create the select sql query
	rows := QueryRowForType(db, sqlStatement) // execute the sql statement
	users = models.ParseUserArray(rows)       // build array
	defer rows.Close()                        // rows closing
	return users                              // return users
}

// update user in the DB
func updateUser(id int64, user models.User, db *sql.DB) int64 {
	PingOrPanic(db)                                                                                             // check the connection
	sqlStatement := `UPDATE users SET name=$2, email=$3, password=$4 WHERE userId=$1`                           // create the update sql query
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id, user.Name, user.Email, user.Password) // check how many rows affected
	return rowsAffected
}

// delete user in the DB
func deleteUser(id int64, db *sql.DB) int64 {
	PingOrPanic(db)                                                       // check the connection
	sqlStatement := `DELETE FROM users WHERE userId=$1`                   // create the delete sql query
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id) // check how many rows affected
	return rowsAffected
}
func insertUser(user models.User, db *sql.DB) int64 {
	var id int64                                   // int for storarge
	canInsert := true                              // can insert yes please
	_, err := getUserByName(string(user.Name), db) // get user by name
	if err != nil {                                //if error = nill
		canInsert = true                     //jank way to use an error
		log.Printf("This user name is free") //log error
	}
	if canInsert == true {
		sqlStatement := `INSERT INTO users (name, password,email,gender,token) VALUES ($1, $2, $3,$4,$5) RETURNING userId` //statement
		err := db.QueryRow(sqlStatement, user.Name, user.Password, user.Email, user.Gender, user.Token).Scan(&id)          //query row
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
	}
	return id
}

//give this a name and password and it spits out an api response with a token
func checkLogin(name string, password string, db *sql.DB) models.APIResponse {
	PingOrPanic(db)                                                         // check the connection
	sqlStatement := `SELECT * FROM users WHERE name = $1 and password = $2` // build the statements
	rows := QueryRowForType(db, sqlStatement, name, password)               // execute the sql statement
	var users []models.User                                                 // build the users array
	users = models.ParseUserArray(rows)                                     // populate the users array
	if len(users) == 0 {                                                    // if no users match this
		return models.BuildAPIResponseFail("No users founds.", nil)
	}
	token := GenerateToken(users[0])                                // populate the response
	user := users[0]                                                // take the first user with this info
	user.Token = token                                              // set the token
	api := models.BuildAPIResponseSuccess("Login Successful", user) // send a user response
	defer rows.Close()
	return api
}
