package web

import (
	"database/sql"
	b64 "encoding/base64"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//QueryRowForID query a statement for a specific id return
func QueryRowForID(db *sql.DB, sqlStatement string, args ...interface{}) int64 {
	var id int64
	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, args).Scan(&id)
	//TODO make this error message less bad
	if err != nil {
		log.Printf("Unable to execute the query. %v\n", err)
		log.Printf("this was likely and insert")

	}
	fmt.Printf("Inserted a single record %v\n", id)

	// return the inserted id
	return id
}

//QueryRowForType query for any type. returns the rows to be worked with. close the rows when your done
func QueryRowForType(db *sql.DB, sqlStatement string, args ...interface{}) *sql.Rows {
	// execute the sql statement

	if len(args) == 0 {
		row, err := db.Query(sqlStatement)
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		return row
	} else if len(args) == 1 {
		row, err := db.Query(sqlStatement, args[0])
		fmt.Println(args[0])
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		return row
	} else {
		//log.Printf("Doing multiArgQuery")
		if len(args) > 2 {
			panic("Oh god. to many args in Query for row type")
		}
		row, err := db.Query(sqlStatement, args[0], args[1])
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		return row
	}
}

//PingOrPanic this is to ping the database and panic the error if something is wrong
func PingOrPanic(db *sql.DB) {
	// check the connection
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

//Execute does statemetns
//TODO maybe make this return a response
func Execute(db *sql.DB, sqlStatement string) {
	//execute the sql statement and return a response
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
}

//GenerateToken take a user and use the stuff to make base64 encoded login token. DO NOT DO THIS IN PRODUCTION
func GenerateToken(user models.User) string {
	data := user.Name + user.Password
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc
}

//ExecStatementAndGetRowsAffected Execute an sql statement and return an int of the rows affected while handling all errors
func ExecStatementAndGetRowsAffected(db *sql.DB, sqlStatement string, args ...interface{}) int64 {

	// execute the sql statement
	res, err := db.Exec(sqlStatement, args...)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v\n", rowsAffected)

	return rowsAffected
}
