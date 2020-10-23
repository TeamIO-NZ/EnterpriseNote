package web

import (
	"database/sql"
	"fmt"
	"log"
)

//------------------------------Main Request Handler functions--------------------------------//

//getNote
func getNote(id int64) (Note, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var note Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&note.ID, &note.Title, &note.Desc, &note.Content)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return note, nil
	case nil:
		return note, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return note, err
}

// get one user from the DB by its userid
func getAllNotes() ([]Note, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var notes []Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var note Note

		// unmarshal the row object to user
		err = rows.Scan(&note.ID, &note.Title, &note.Desc, &note.Content)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		notes = append(notes, note)

	}

	// return empty user on error
	return notes, err
}

// update user in the DB
func updateNote(id int64, note Note) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE notes SET title=$2, description=$3, contents=$4 WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, note.Title, note.Desc, note.Content)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func deleteNote(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM notes WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

//insert note
func insertNote(note Note) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO notes (title, description, contents) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, note.Title, note.Desc, note.Title).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}
