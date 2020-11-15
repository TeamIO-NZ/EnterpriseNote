package web

import (
	"database/sql"
	"fmt"
	"log"

	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//------------------------------Main Request Handler functions- Users-------------------------------//

//getNote
// get one user from the DB by its userid
func getUser(id int64, db *sql.DB) models.User {
	PingOrPanic(db)

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE userId=$1`

	// execute the sql statement
	row := QueryRowForType(db, sqlStatement, id)

	// unmarshal the row object to user
	user, _ := models.ParseSingleUser(row)
	defer row.Close()
	// return empty user on error
	return user
}

//getNote
// get one user from the DB by its userid
func getUserByName(name string, db *sql.DB) (models.User, error) {
	// check the connection
	PingOrPanic(db)
	//fmt.Println("searching user by name")
	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE name=$1`

	// execute the sql statement
	row := QueryRowForType(db, sqlStatement, name)

	// unmarshal the row object to user
	user, err := models.ParseSingleUser(row)
	defer row.Close()
	// return empty user on error
	return user, err
}

// get one user from the DB by its userid
func getUserByEmail(email string, db *sql.DB) (models.User, error) {
	// check the connection
	PingOrPanic(db)
	//fmt.Println("searching user by name")
	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE email=$1`

	// execute the sql statement
	row := QueryRowForType(db, sqlStatement, email)

	// unmarshal the row object to user
	user, err := models.ParseSingleUser(row)
	defer row.Close()
	// return empty user on error
	return user, err
}

// get one user from the DB by its userid
func getAllUsers(db *sql.DB) []models.User {
	// check the connection
	PingOrPanic(db)

	var users []models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows := QueryRowForType(db, sqlStatement)

	users = models.ParseUserArray(rows)
	// // iterate over the rows
	// for rows.Next() {
	// 	user, _ := models.ParseSingleUser(rows)
	// 	fmt.Println(user.Name)
	// 	// append the user in the users slice
	// 	users = append(users, user)
	// }
	// return empty user on error
	defer rows.Close()
	return users
}

// update user in the DB
func updateUser(id int64, user models.User, db *sql.DB) int64 {

	// check the connection
	PingOrPanic(db)

	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, email=$3, password=$4 WHERE userId=$1`

	// check how many rows affected
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id, user.Name, user.Email, user.Password)

	return rowsAffected
}

// delete user in the DB
func deleteUser(id int64, db *sql.DB) int64 {

	// check the connection
	PingOrPanic(db)

	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE userId=$1`

	// check how many rows affected
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id)

	return rowsAffected
}

//insert note into the database
func insertUser(user models.User, db *sql.DB) int64 {
	// check the connection
	PingOrPanic(db)

	var id int64
	canInsert := false
	_, err := getUserByName(string(user.Name), db)
	if err != nil {
		canInsert = true
		log.Printf("This user name is free")
	}

	//if you can insert the user then do so
	if canInsert == true {
		sqlStatement := `INSERT INTO users (name, email, password, gender) VALUES ($1, $2, $3,$4) RETURNING userId`
		id = QueryRowForID(db, sqlStatement, &user.Name, &user.Email, &user.Password, &user.Gender)
		fmt.Printf("Created user with id of %d", id)
	}
	return id
}
func testInsert(user models.User, db *sql.DB) int64 {
	//fmt.Println(user.Name)
	var id int64
	canInsert := true
	_, err := getUserByName(string(user.Name), db)
	if err != nil {
		canInsert = true
		log.Printf("This user name is free")
	}
	if canInsert == true {
		sqlStatement := `INSERT INTO users (name, password,email,gender,token) VALUES ($1, $2, $3,$4,$5) RETURNING userId`
		//fmt.Printf("offending id = %d", )
		err := db.QueryRow(sqlStatement, user.Name, user.Password, user.Email, user.Gender, user.Token).Scan(&id)
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		fmt.Printf("Inserted a single record %v\n", id)
	}
	return id
}

//give this a name and password and it spits out an api response with a token
func checkLogin(name string, password string, db *sql.DB) models.APIResponse {
	// check the connection
	PingOrPanic(db)
	//build the statements
	sqlStatement := `SELECT * FROM users WHERE name = $1 and password = $2`
	// execute the sql statement
	rows := QueryRowForType(db, sqlStatement, name, password)
	//build the users array
	var users []models.User
	//populate the users array
	users = models.ParseUserArray(rows)
	if len(users) == 0 {
		return models.BuildAPIResponseFail("No users founds.", nil)
	}
	//populate the response
	token := GenerateToken(users[0])
	user := users[0]
	user.Token = token
	fmt.Printf(token)
	api := models.BuildAPIResponseSuccess("Login Successful", user)
	defer rows.Close()
	return api
}
