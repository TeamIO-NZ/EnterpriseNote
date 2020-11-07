package models

import (
	b64 "encoding/base64"
	"fmt"
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

//DecodeDumbStringsToArrays decode the array to a string
//TODO ignore my jank error discarding
func DecodeDumbStringsToArrays(dumbString []byte) []int {
	data := string(dumbString)
	sEnc, _ := b64.StdEncoding.DecodeString(data)
	s := strings.Split(string(sEnc), "{,}")
	fmt.Println(s)
	var v []int
	// for i := 0; i < len(s); i++ {
	// 	b := strconv.Itoa(s[i])
	// 	v = append(v)
	// }
	return v
}
