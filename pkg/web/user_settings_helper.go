package web

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

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
