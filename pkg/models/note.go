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
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
	Owner   int    `json:"Owner"`
	Viewers []int  `json:"Viewers"`
	Editors []int  `json:"Editors"`
}

//ParseStringForArrayNumbers Break a string into an array of numbers
func ParseStringForArrayNumbers(stringToBreak string, array []int) (arrayToReturn []int) {
	stringToBreak = strings.Replace(stringToBreak, "}", "", -1)
	stringToBreak = strings.Replace(stringToBreak, "{", "", -1)

	//fmt.Println(viewers)
	//split the string by the comma
	split := strings.Split(stringToBreak, ",")
	//loop the resulting array and convert every item to a number
	for n := range split {
		str, err := strconv.Atoi(fmt.Sprint(n))
		//crash if its not a number
		if err != nil {
			log.Println("Oh god its broken")
		}
		//append it to the array
		array = append(array, str)
	}
	return array
}
func ParseNoteArray(rows *sql.Rows) []Note {
	var note Note
	var notes []Note
	for rows.Next() {
		var viewers string
		var editors string
		fmt.Println("scanning a row of note stuff. awaiting crash")
		// unmarshal the row object to user
		err := rows.Scan(&note.ID, &note.Title, &note.Desc, &note.Content, &note.Owner, &viewers, &editors)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		note.Viewers = ParseStringForArrayNumbers(viewers, note.Viewers)
		note.Editors = ParseStringForArrayNumbers(editors, note.Editors)
		// // append the user in the users slice
		notes = append(notes, note)
	}
	return notes
}
