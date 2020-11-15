package web

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//------------------------------Main Request Handler functions- Notes-------------------------------//

//getNote
// get one user from the DB by its userid
func getNote(id int64, db *sql.DB) models.Note {

	PingOrPanic(db)
	// create a user of models.User type
	var note models.Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes WHERE id=$1`

	// execute the sql statement
	row := QueryRowForType(db, sqlStatement, id)
	//scan the note properly
	note = models.ParseSingleNote(row)

	// return empty user on error
	return note
}

// get one user from the DB by its userid
func getAllNotes(db *sql.DB) []models.Note {

	PingOrPanic(db)

	var notes []models.Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes`

	// execute the sql statement
	row := QueryRowForType(db, sqlStatement)
	//close the statement
	defer row.Close()

	// append the user in the users slice
	notes = models.ParseNoteArray(row)

	return notes
}

// update user in the DB
func updateNote(id int64, note models.Note, db *sql.DB) int64 {

	PingOrPanic(db)
	// create the update sql query
	sqlStatement := `UPDATE notes SET id=$1, title=$2, description=$3, contents=$4, owner=$5,viewers=$6,editors=$7 WHERE id=$1`

	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, note.ID, note.Title, note.Desc, note.Content, note.Owner, pq.Array(note.Viewers), pq.Array(note.Editors))

	return rowsAffected
}

// delete user in the DB
func deleteNote(id int, db *sql.DB) int64 {

	PingOrPanic(db)

	// // create the delete sql query
	sqlStatement := `DELETE FROM notes WHERE id=$1`

	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id)

	return rowsAffected
}

//insert note into the database
func insertNote(note models.Note, db *sql.DB) (int64, error) {
	PingOrPanic(db)
	var id int64
	// create the insert sql query
	// returning id will return the id of the inserted note
	sqlStatement := `INSERT INTO notes (title, description, contents, owner, viewers, editors) VALUES ($1,$2, $3,$4,$5,$6) RETURNING id`
	// the inserted id will store in this id
	res, err := db.Exec(sqlStatement, note.Title, note.Desc, note.Content, note.Owner, pq.Array(note.Viewers), pq.Array(note.Editors))
	if err != nil {
		log.Printf("note: %s is the offending note", note.Title)
		log.Printf("Unable to execute the query. %v\n", err)
	}
	fmt.Println(res)
	return id, err
}

//------------------------------User Common Request Handler functions- Notes-------------------------------//

//getNote
// get one user from the DB by its userid
func getUserSettings(id int64, db *sql.DB) models.UserSettings {

	PingOrPanic(db)
	// create a user of models.User type
	var userSettings models.UserSettings

	// create the select sql query
	sqlStatement := `SELECT * FROM usersettings WHERE id=$1`

	// execute the sql statement
	row := QueryRowForType(db, sqlStatement, id)
	//scan the note properly
	userSettings = models.ParseSingleUserSetting(row)
	defer row.Close()
	// return empty user on error
	return userSettings
}

// update user in the DB
func updateuserSettings(userSettings models.UserSettings, db *sql.DB) (int64, error) {

	PingOrPanic(db)
	// create the update sql query
	sqlStatement := `UPDATE usersettings SET id=$1, viewers=$2, editors=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement, userSettings.ID, pq.Array(userSettings.Viewers), pq.Array(userSettings.Editors))
	if err != nil {
		log.Printf("User setting not updated due to: %v", err)
	}
	return res.RowsAffected()
}

// delete user in the DB
func deleteUserSettings(id int64, db *sql.DB) int64 {

	PingOrPanic(db)

	// // create the delete sql query
	sqlStatement := `DELETE FROM usersettings WHERE id=$1`

	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id)

	return rowsAffected
}

//insert note into the database
func insertUserSettings(userSettings models.UserSettings, db *sql.DB) (int64, error) {
	PingOrPanic(db)
	var id int64
	// create the insert sql query
	// returning id will return the id of the inserted note
	sqlStatement := `INSERT INTO usersettings (id, viewers, editors) VALUES ($1,$2,$3) RETURNING id`
	// the inserted id will store in this id
	err := db.QueryRow(sqlStatement, userSettings.ID, pq.Array(userSettings.Viewers), pq.Array(userSettings.Editors)).Scan(&id)
	if err != nil {
		log.Printf("note: %v is the offending note", userSettings.Editors)
		log.Printf("Unable to execute the query. %v\n", err)
	}
	log.Printf("id: %d", id)
	return id, err
}

//insert note into the database
// func insertUserSettings(userSettings models.UserSettings, db *sql.DB) (id int64) {
// 	PingOrPanic(db)
// 	// create the insert sql query
// 	// returning id will return the id of the inserted note
// 	sqlStatement := `INSERT INTO usersettings (viewers, editors) VALUES ($1, $2) RETURNING id`
// 	// the inserted id will store in this id
// 	err := db.QueryRow(sqlStatement, pq.Array(userSettings.Viewers), pq.Array(userSettings.Editors)).Scan(&id)
// 	if err != nil {
// 		log.Printf("Error inserting statement")
// 		log.Printf("Offending statement was %s", sqlStatement)
// 	}
// 	log.Printf("Id is: %d", id)
// 	// return the inserted id
// 	return id
// }

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

/*------------------specific searches-----------**/

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

// get one user from the DB by its userid
func getAllNotesUserHasAccessTo(id int, db *sql.DB) []models.Note {
	// check the connection
	PingOrPanic(db)

	var notes []models.Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes 
	WHERE editors @> ARRAY[$1]::int[] 	or viewers @> ARRAY[$1]::int[] or owner = $1`
	row := QueryRowForType(db, sqlStatement, id)
	// iterate over the rows
	notes = models.ParseNoteArray(row)
	defer row.Close()
	// return empty user on error
	return notes
}
