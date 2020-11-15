package web

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//------------------------------Main Request Handler functions- Notes-------------------------------//

//get one note from the DB by its userid
func getNote(id int64, db *sql.DB) models.Note {
	PingOrPanic(db)                                   //error checking for if the database connection exists
	var note models.Note                              // create a user of models.User type
	sqlStatement := `SELECT * FROM notes WHERE id=$1` // create the select sql query
	row := QueryRowForType(db, sqlStatement, id)      // execute the sql statement
	note = models.ParseSingleNote(row)                // scan the note properly
	return note                                       // return empty user on error
}

// get all the notes
func getAllNotes(db *sql.DB) []models.Note {
	PingOrPanic(db)                          //error checking for if the database connection exists
	var notes []models.Note                  //make a note array
	sqlStatement := `SELECT * FROM notes`    // create the select sql query
	row := QueryRowForType(db, sqlStatement) // execute the sql statement
	defer row.Close()                        // close the statement
	notes = models.ParseNoteArray(row)       // append the user in the users slice
	return notes
}

// update note in the DB with the json sent in
func updateNote(id int64, note models.Note, db *sql.DB) int64 {
	PingOrPanic(db)                                                                                                             //error checking for if the database connection exists
	sqlStatement := `UPDATE notes SET id=$1, title=$2, description=$3, contents=$4, owner=$5,viewers=$6,editors=$7 WHERE id=$1` // create the update sql query
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, note.ID, note.Title, note.Desc, note.Content, note.Owner, pq.Array(note.Viewers), pq.Array(note.Editors))
	return rowsAffected
}

// delete user in the DB decided by id
func deleteNote(id int, db *sql.DB) int64 {

	PingOrPanic(db)                                                       //error checking for if the database connection exists
	sqlStatement := `DELETE FROM notes WHERE id=$1`                       // create the delete sql query
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id) //execute the statement
	return rowsAffected
}

//insert note into the database using the json object provided
func insertNote(note models.Note, db *sql.DB) int64 {
	PingOrPanic(db)                                                                                                                      //error checking for if the database connection exists
	var id int64                                                                                                                         // returning id will return the id of the inserted note
	sqlStatement := `INSERT INTO notes (title, description, contents, owner, viewers, editors) VALUES ($1,$2, $3,$4,$5,$6) RETURNING id` // create the insert sql query
	_, err := db.Exec(sqlStatement, note.Title, note.Desc, note.Content, note.Owner, pq.Array(note.Viewers), pq.Array(note.Editors))
	if err != nil {
		log.Printf("note: %s is the offending note", note.Title)
		log.Printf("Unable to execute the query. %v\n", err)
	}
	return id
}

// get one user from the DB by its userid
func getAllNotesUserHasAccessTo(id int, db *sql.DB) []models.Note {
	PingOrPanic(db)         // check the connection
	var notes []models.Note // create a note array
	sqlStatement := `SELECT * FROM notes  
	WHERE editors @> ARRAY[$1]::int[] or 
	viewers @> ARRAY[$1]::int[] or owner = $1` // create the select sql query

	row := QueryRowForType(db, sqlStatement, id) // iterate over the rows
	notes = models.ParseNoteArray(row)           // make the notes
	defer row.Close()
	return notes
}
