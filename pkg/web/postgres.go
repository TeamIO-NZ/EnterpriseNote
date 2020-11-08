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
	sqlStatement := `UPDATE notes SET title=$2, description=$3, contents=$4, owner=$5,viewers=$6,editors=$7  WHERE id=$1`

	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id, note.Title, note.Desc, note.Content, &note.Owner, pq.Array(&note.Viewers), pq.Array(&note.Editors))

	return rowsAffected
}

// delete user in the DB
func deleteNote(id int64, db *sql.DB) int64 {

	PingOrPanic(db)

	// // create the delete sql query
	sqlStatement := `DELETE FROM notes WHERE id=$1`

	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id)

	return rowsAffected
}

//insert note into the database
func insertNote(note models.Note, db *sql.DB) int64 {
	PingOrPanic(db)
	// create the insert sql query
	// returning id will return the id of the inserted note
	sqlStatement := `INSERT INTO notes (title, description, contents, owner, viewers, editors) VALUES ($1, $2, $3,$4,$5,$6) RETURNING id`
	// the inserted id will store in this id
	id := QueryRowForID(db, sqlStatement, note.Title, note.Desc, note.Title, note.Owner, pq.Array(note.Viewers), pq.Array(note.Editors))

	// return the inserted id
	return id
}

//------------------------------Main Request Handler functions- Users-------------------------------//

//getNote
// get one user from the DB by its userid
func getUser(id int64, db *sql.DB) models.User {
	PingOrPanic(db)

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE id=$1`

	// execute the sql statement
	row := QueryRowForType(db, sqlStatement, id)

	// unmarshal the row object to user
	user := models.ParseSingleUser(row)

	// return empty user on error
	return user
}

//getNote
// get one user from the DB by its userid
func getUserByName(name string, db *sql.DB) models.User {
	// check the connection
	PingOrPanic(db)
	//fmt.Println("searching user by name")
	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE name=$1`

	// execute the sql statement
	row := QueryRowForType(db, sqlStatement, name)

	// unmarshal the row object to user
	user := models.ParseSingleUser(row)

	// return empty user on error
	return user
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

	// iterate over the rows
	for rows.Next() {
		user := models.ParseSingleUser(rows)
		// append the user in the users slice
		users = append(users, user)
	}
	// return empty user on error
	return users
}

// update user in the DB
func updateUser(id int64, user models.User, db *sql.DB) int64 {

	// check the connection
	PingOrPanic(db)

	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, email=$3, password=$4 WHERE id=$1`

	// check how many rows affected
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id, user.Name, user.Email, user.Password)

	return rowsAffected
}

// delete user in the DB
func deleteUser(id int64, db *sql.DB) int64 {

	// check the connection
	PingOrPanic(db)

	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE id=$1`

	// check how many rows affected
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id)

	return rowsAffected
}

//insert note into the database
func insertUser(user models.User, db *sql.DB) int64 {
	// check the connection
	PingOrPanic(db)

	var id int64
	canInsert := true
	u := getUserByName(string(user.Name), db)
	//check to see if a user name is taken
	if u.Name == user.Name {
		canInsert = false
		log.Printf("This user name is already taken\n")
	}
	//if you can insert the user then do so
	if canInsert == true {
		sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
		id = QueryRowForID(db, sqlStatement, &user.Name, &user.Password, &user.Email)
	}
	return id
}

/*------------------specific searches-----------**/
// get one user from the DB by its userid
func getSpecificNotes(db *sql.DB, id int, targetFunction string, prefix string) []models.Note {
	// check the connection
	PingOrPanic(db)

	var notes []models.Note

	// create the select sql query
	//TODO work out this statement
	sqlStatement := ` `
	fmt.Println(id)
	fmt.Println(targetFunction)
	fmt.Println(prefix)
	switch targetFunction {
	case "1":
		{
			//TODO a sentence with a given prefix and/or suffix.
			sqlStatement = `SELECT * FROM notes 
			WHERE editors @> ARRAY[$1]::int[] or viewers @> ARRAY[$1]::int[] or owner = $1
			AND contents LIKE $2`
			break
		}
	case "2":
		{
			//TODO-a phone number with a given area code and optionally a consecutive sequence of numbers that are part 0f that number.
			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	case "3":
		{
			//TODO an email address on a domain that is only partially provided.
			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	case "4":
		{
			//TODO text that contains at least three of the following case-insensitive words: meeting, minutes, agenda, action, attendees, apologies.

			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	case "5":
		{
			//TODO a word in all capitals of three characters or more.
			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	}
	sqlStatement = `SELECT * FROM notes 
			WHERE editors @> ARRAY[$1]::int[] or viewers @> ARRAY[$1]::int[] or owner = $1`
	//AND contents LIKE $2`
	rows := QueryRowForType(db, sqlStatement, id)
	// iterate over the rows
	notes = models.ParseNoteArray(rows)
	// close the statement
	defer rows.Close()
	return notes
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
	//populate the response

	api := models.BuildAPIResponseSuccess("Login Successful", GenerateToken(users[0]))
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
