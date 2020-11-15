package web

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//CreateTable populates database with fake data
func CreateTable(db *sql.DB) {
	PingOrPanic(db)
	//prepares to close database when done
	defer db.Close()

	//create the base notes table for if it doesn't exist
	sqlStatement := `DROP TABLE IF EXISTS notes;`
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP TABLE IF EXISTS users;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS user_id_sequence;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS notes_userid_sequence;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS usersettings_id_sequence;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP SEQUENCE IF EXISTS zero_index_auto_increment;`
	//execute the sql statement and return a response
	Execute(db, sqlStatement)
	//create the base notes table for if it doesn't exist
	sqlStatement = `DROP TABLE IF EXISTS userSettings;`
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS userSettings (
			id Serial PRIMARY KEY,
			viewers integer[],
			editors integer[]
		);`
	Execute(db, sqlStatement)
	UserSettings := []models.UserSettings{
		{
			ID:      1,
			Viewers: []int{6, 2, 3},
			Editors: []int{4, 5},
		},
		{
			ID:      2,
			Viewers: []int{1, 2, 3},
			Editors: []int{4, 5},
		},
		{
			ID:      3,
			Viewers: []int{1, 2, 3},
			Editors: []int{4, 5},
		},
	}
	for _, userSettings := range UserSettings {
		fmt.Println(userSettings.ID)
		var id int64
		sqlStatement := `INSERT INTO userSettings (id, viewers,editors) VALUES ($1,$2, $3) RETURNING id`

		err := db.QueryRow(sqlStatement, userSettings.ID, pq.Array(userSettings.Viewers), pq.Array(userSettings.Editors)).Scan(&id)
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		fmt.Printf("Inserted a single record %v\n", id)

	}
	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS users (
		userId serial PRIMARY KEY,
		name TEXT,
		password TEXT,
		gender TEXT,
		email TEXT,
		token TEXT,
		userSettingsId int default 1
	);`
	Execute(db, sqlStatement)

	//create the base notes table for if it doesn't exist
	sqlStatement = `CREATE TABLE IF NOT EXISTS notes (
			id serial PRIMARY KEY,
			title TEXT,
			description TEXT,
			contents TEXT,
			owner INT,
			viewers integer[],
			editors integer[],
			FOREIGN KEY (owner)	REFERENCES users (userId) on delete cascade on update cascade
		);`
	Execute(db, sqlStatement)

	// Execute(db, sqlStatement)
	Users := []models.User{
		{
			Name:           "lithial",
			Password:       "1234",
			Email:          "me@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "joe",
			Password:       "1234",
			Email:          "you@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "peter",
			Password:       "1234",
			Email:          "us@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "arran",
			Password:       "1234",
			Email:          "re@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "finn",
			Password:       "1234",
			Email:          "de@james.me",
			UserSettingsID: 0,
		},
		{
			Name:           "sam",
			Password:       "1234",
			Email:          "la@james.me",
			UserSettingsID: 0,
		},
	}

	for _, user := range Users {
		//fmt.Println(user.Name)
		var id int64
		canInsert := true
		u, _ := getUserByName(string(user.Name), db)
		if u.Name == user.Name {
			canInsert = false
			log.Printf("This user name is already taken\n")
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
	}
}
