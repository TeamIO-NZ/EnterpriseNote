package web

import (
	"database/sql"
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

	// create the insert sql query
	// returning id will return the id of the inserted note
	sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	// the inserted id will store in this id
	id := QueryRowForID(db, sqlStatement, &user.ID, &user.Name, &user.Password, &user.Email, &user.Token)
	return id
}

/*------------------specific searches-----------**/
// get one user from the DB by its userid
func getSpecificNotes(searchType int, db *sql.DB) ([]models.Note, error) {
	// check the connection
	PingOrPanic(db)

	var notes []models.Note

	// create the select sql query
	//TODO work out this statement
	sqlStatement := ` `

	switch searchType {
	case 1:
		{
			//TODO a sentence with a given prefix and/or suffix.
			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	case 2:
		{
			//TODO-a phone number with a given area code and optionally a consecutive sequence of numbers that are part 0f that number.
			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	case 3:
		{
			//TODO an email address on a domain that is only partially provided.
			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	case 4:
		{
			//TODO text that contains at least three of the following case-insensitive words: meeting, minutes, agenda, action, attendees, apologies.

			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	case 5:
		{
			//TODO a word in all capitals of three characters or more.
			sqlStatement = `SELECT * FROM notes WHERE contents LIKE ''`
			break
		}
	}

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var note models.Note

		// unmarshal the row object to user
		err = rows.Scan(&note.ID, &note.Title, &note.Desc, &note.Content)

		if err != nil {
			log.Printf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		notes = append(notes, note)

	}
	return notes, err
}

//give this a name and password and it spits out an api response with a token
func checkLogin(name string, password string, db *sql.DB) models.APIResponse {
	// check the connection
	PingOrPanic(db)

	sqlStatement := `SELECT * FROM users WHERE name = $1 and password = $2`
	// execute the sql statement
	rows := QueryRowForType(db, sqlStatement, name, password)
	var users []models.User

	//TODO theoretically this could return more than one user
	// for rows.Next() {
	// 	user := models.ParseSingleUser(rows)
	// 	// append the user in the users slice
	// 	users = append(users, user)
	// 	numOfUsers := len(users)
	// 	if numOfUsers > 1 {
	// 		log.Printf("More than one user with this name|password combo. please investigate")
	// 	}

	// }
	users = models.ParseUserArray(rows)
	//populate the response
	//TODO implement checks for multiple users with the same user and password
	var api models.APIResponse
	api.Code = 200
	api.Message = "Successful User Acquired"
	api.Data = GenerateToken(users[0]) //use only the first one
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
	WHERE 
	editors @> ARRAY[$1]::int[]
	or viewers @> ARRAY[$1]::int[]
	or owner = $1`
	row := QueryRowForType(db, sqlStatement, id)
	// iterate over the rows
	notes = models.ParseNoteArray(row)
	defer row.Close()
	// return empty user on error
	return notes
}
