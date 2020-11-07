package web

import (
	"database/sql"
	b64 "encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//------------------------------Main Request Handler functions- Notes-------------------------------//

//getNote
// get one user from the DB by its userid
func getNote(id int64) (models.Note, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var note models.Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&note.ID, &note.Title, &note.Desc, &note.Content, &note.Owner, &note.Viewers, &note.Editors)

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
func getAllNotes() ([]models.Note, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var notes []models.Note

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
		var note models.Note

		// unmarshal the row object to user
		err = rows.Scan(&note.ID, &note.Title, &note.Desc, &note.Content, &note.Owner, &note.Viewers, &note.Editors)

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
func updateNote(id int64, note models.Note) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE notes SET title=$2, description=$3, contents=$4, owner=$5,viewers=$6,editors=$7  WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, note.Title, note.Desc, note.Content, &note.Owner, pq.Array(&note.Viewers), pq.Array(&note.Editors))

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

//insert note into the database
func insertNote(note models.Note) int64 {
	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()
	// create the insert sql query
	// returning id will return the id of the inserted note
	sqlStatement := `INSERT INTO notes (title, description, contents, owner, viewers, editors) VALUES ($1, $2, $3,$4,$5,$6) RETURNING id`
	// the inserted id will store in this id
	var id int64
	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, note.Title, note.Desc, note.Title, note.Owner, pq.Array(note.Viewers), pq.Array(note.Editors)).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

//------------------------------Main Request Handler functions- Users-------------------------------//

//getNote
// get one user from the DB by its userid
func getUser(id int64) (models.User, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var user models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

// get one user from the DB by its userid
func getAllUsers() ([]models.User, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var users []models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)

	}

	// return empty user on error
	return users, err
}

// update user in the DB
func updateUser(id int64, user models.User) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, email=$3, password=$4 WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, user.Name, user.Email, user.Password)

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
func deleteUser(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE id=$1`

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

//insert note into the database
func insertUser(user models.User) int64 {
	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()
	// create the insert sql query
	// returning id will return the id of the inserted note
	sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	// the inserted id will store in this id
	var id int64
	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

/*------------------specific searches-----------**/
// get one user from the DB by its userid
func getSpecificNotes(searchType int) ([]models.Note, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

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
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var note models.Note

		// unmarshal the row object to user
		err = rows.Scan(&note.ID, &note.Title, &note.Desc, &note.Content)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		notes = append(notes, note)

	}
	return notes, err
}

//give this a name and password and it spits out an api response with a token
func checkLogin(name string, password string) (models.APIResponse, error) {
	db := createConnection()
	// close the db connection
	defer db.Close()
	sqlStatement := `SELECT * FROM users WHERE name = $1 and password = $2`
	// execute the sql statement
	rows, err := db.Query(sqlStatement, name, password)
	var users []models.User
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	//TODO theoretically this could return more than one user
	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Token)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)
	}
	// close the statement
	//populate the response
	//TODO implement checks for multiple users with the same user and password
	var api models.APIResponse
	api.Code = 200
	api.Message = "Successful User Acquired"
	api.Data = generateToken(users[0]) //use only the first one
	defer rows.Close()
	return api, err
}

//take a user and use the stuff to make base64 encoded login token. DO NOT DO THIS IN PRODUCTION
func generateToken(user models.User) string {
	data := user.Name + user.Password
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc
}

// get one user from the DB by its userid
func getAllNotesUserHasAccessTo(id int) ([]models.Note, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var notes []models.Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes 
	WHERE 
	editors @> ARRAY[$1]::int[]
	or viewers @> ARRAY[$1]::int[]
	or owner = $1`
	//select * from users where id = ANY(ARRAY [1,2])
	// execute the sql statement
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var note models.Note

		var viewers string
		var editors string
		// unmarshal the row object to user
		err = rows.Scan(&note.ID, &note.Title, &note.Desc, &note.Content, &note.Owner, &viewers, &editors)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		//convert the weird string to an int array
		// fmt.Println(viewers)
		// fmt.Println(editors)
		//remove the curly brackets without my brain hating regex to much
		viewers = strings.Replace(viewers, "{", "", -1)
		viewers = strings.Replace(viewers, "}", "", -1)
		//fmt.Println(viewers)
		//split the string by the comma
		v := strings.Split(viewers, ",")
		//loop the resulting array and convert every item to a number
		for n := range v {
			str, err := strconv.Atoi(fmt.Sprint(n))
			//crash if its not a number
			if err != nil {
				log.Println("Oh god its broken")
			}
			//append it to the array
			note.Viewers = append(note.Viewers, str)
		}
		//TODO make this a function because i do it twice
		editors = strings.Replace(viewers, "{", "", -1)
		editors = strings.Replace(viewers, "}", "", -1)
		fmt.Println(editors)
		e := strings.Split(editors, ",")
		for n := range e {
			str, err := strconv.Atoi(fmt.Sprint(n))
			if err != nil {
				log.Println("Oh god its broken")
			}
			note.Editors = append(note.Editors, str)

		}
		// append the user in the users slice
		notes = append(notes, note)

	}

	// return empty user on error
	return notes, err
}
