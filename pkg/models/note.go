package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//Note a note object for json
type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Owner   int    `json:"owner"`
	Viewers []int  `json:"viewers"`
	Editors []int  `json:"editors"`
}

//ParseStringForArrayNumbers Break a string into an array of numbers
func ParseStringForArrayNumbers(stringToBreak string) (arrayToReturn []int) {
	stringToBreak = strings.Replace(stringToBreak, "}", "", -1)
	stringToBreak = strings.Replace(stringToBreak, "{", "", -1)

	//split the string by the comma
	split := strings.Split(stringToBreak, ",")
	splitLen := len(split)
	if splitLen != 0 {
		for i := 0; i < splitLen; i++ {
			str, err := strconv.Atoi(split[i])
			//crash if its not a number
			if err != nil {
				log.Println("Oh god its broken")
			}
			//append it to the array
			arrayToReturn = append(arrayToReturn, str)
		}
		//loop the resulting array and convert every item to a number
		return arrayToReturn
	}
	return arrayToReturn
}

//ParseNoteArray parses a note array and returns the entire array
func ParseNoteArray(rows *sql.Rows) []Note {
	var note Note
	var notes []Note
	// fmt.Println("Begin scanning rows")
	// fmt.Printf("row status = %t\n", rows.Next())
	for rows.Next() {
		var viewers string = ""
		var editors string = ""
		//fmt.Println("scanning a row of note stuff. awaiting crash")
		// unmarshal the row object to user
		err := rows.Scan(&note.ID, &note.Title, &note.Desc, &note.Content, &note.Owner, &viewers, &editors)
		if err != nil {
			log.Printf("ParseSingleNote: Unable to scan the row. %v\n", err)
		}

		note.Viewers = ParseStringForArrayNumbers(viewers)
		note.Editors = ParseStringForArrayNumbers(editors)
		// // append the user in the users slice
		notes = append(notes, note)
	}
	return notes
}

//ParseSingleNote Parses a single note and returns it
func ParseSingleNote(row *sql.Rows) Note {
	var note Note
	var viewers string
	var editors string
	fmt.Println("scanning a row of note stuff. awaiting crash")
	// unmarshal the row object to user
	if row.Next() {
		err := row.Scan(&note.ID, &note.Title, &note.Desc, &note.Content, &note.Owner, &viewers, &editors)
		if err != nil {
			log.Printf("ParseSingleNote: Unable to scan the row. %v\n", err)
		}
		note.Viewers = ParseStringForArrayNumbers(viewers)
		note.Editors = ParseStringForArrayNumbers(editors)
	}
	//return the note
	return note
}
