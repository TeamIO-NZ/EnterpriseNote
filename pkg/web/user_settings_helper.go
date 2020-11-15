package web

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//------------------------------User Common Request Handler functions- Notes-------------------------------//

//getNote
// get one user settings from the DB by its userid
func getUserSettings(id int64, db *sql.DB) models.UserSettings {

	PingOrPanic(db)                                          //check database
	var userSettings models.UserSettings                     // create a user of models.User type
	sqlStatement := `SELECT * FROM usersettings WHERE id=$1` // create the select sql query
	row := QueryRowForType(db, sqlStatement, id)             // execute the sql statement
	userSettings = models.ParseSingleUserSetting(row)        //scan the note properly
	defer row.Close()
	return userSettings
}

// update user settings in the DB
func updateuserSettings(userSettings models.UserSettings, db *sql.DB) (int64, error) {

	PingOrPanic(db)                                                                     // check database
	sqlStatement := `UPDATE usersettings SET id=$1, viewers=$2, editors=$3 WHERE id=$1` // create the update sql query
	res, err := db.Exec(sqlStatement, userSettings.ID, pq.Array(userSettings.Viewers), pq.Array(userSettings.Editors))
	if err != nil {
		log.Printf("User setting not updated due to: %v", err)
	}
	return res.RowsAffected()
}

// delete user settings in the DB
func deleteUserSettings(id int64, db *sql.DB) int64 {
	PingOrPanic(db)                                        // check database
	sqlStatement := `DELETE FROM usersettings WHERE id=$1` // create the delete sql query
	rowsAffected := ExecStatementAndGetRowsAffected(db, sqlStatement, id)
	return rowsAffected
}

//insert note into the database
func insertUserSettings(userSettings models.UserSettings, db *sql.DB) (int64, error) {
	PingOrPanic(db)                                                                                                             // check database
	var id int64                                                                                                                // returning id will return the id of the inserted note
	sqlStatement := `INSERT INTO usersettings (id, viewers, editors) VALUES ($1,$2,$3) RETURNING id`                            // create the insert sql query
	err := db.QueryRow(sqlStatement, userSettings.ID, pq.Array(userSettings.Viewers), pq.Array(userSettings.Editors)).Scan(&id) // the inserted id will store in this id
	if err != nil {
		log.Printf("note: %v is the offending note", userSettings.Editors)
		log.Printf("Unable to execute the query. %v\n", err)
	}
	log.Printf("id: %d", id)
	return id, err
}
