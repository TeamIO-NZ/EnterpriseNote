package web

import (
	"database/sql"
	b64 "encoding/base64"
	"log"

	_ "github.com/lib/pq" //its lib pq
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//QueryRowForID query a statement for a specific id return
func QueryRowForID(db *sql.DB, sqlStatement string, args ...interface{}) int64 {
	var id int64                                     // Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, args).Scan(&id) // execute the sql statement
	if err != nil {
		log.Printf("Unable to execute the query. %v\n", err)
	}
	return id // return the inserted id

}

//QueryRowForType query for any type. returns the rows to be worked with. close the rows when your done
func QueryRowForType(db *sql.DB, sqlStatement string, args ...interface{}) *sql.Rows {
	if len(args) == 0 {
		row, err := db.Query(sqlStatement)
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		return row
	} else if len(args) == 1 {
		row, err := db.Query(sqlStatement, args[0])
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		return row
	} else {
		if len(args) > 2 {
			log.Printf("Oh god. to many args in Query for row type\n")
		}
		row, err := db.Query(sqlStatement, args[0], args[1])
		if err != nil {
			log.Printf("Unable to execute the query. %v\n", err)
		}
		return row
	}
}

//PingOrPanic this is to ping the database and panic the error if something is wrong
func PingOrPanic(db *sql.DB) *sql.DB {
	err := db.Ping() // check the connection
	if err != nil {
		log.Printf("Database Error %v", err)
		db.Close()
		newDb, _ := CreateConnection()
		return newDb
	}
	return nil
}

//Execute does statements
func Execute(db *sql.DB, sqlStatement string) {
	_, err := db.Exec(sqlStatement) //execute the sql statement and return a response
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
	res, err := db.Exec(sqlStatement, args...) // execute the sql statement
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected() // check how many rows affected
	if err != nil {
		log.Printf("Error while checking the affected rows. %v", err)
	}
	return rowsAffected
}
