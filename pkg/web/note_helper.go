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

	var note models.Note                              // create a user of models.User type
	sqlStatement := `SELECT * FROM notes WHERE id=$1` // create the select sql query
	row := QueryRowForType(db, sqlStatement, id)      // execute the sql statement
	note = models.ParseSingleNote(row)                // scan the note properly
	return note                                       // return empty user on error
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
func insertNote(note models.Note, db *sql.DB) int64 {
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
	return id
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
