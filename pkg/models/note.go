package models

//Note a note object for json
type Note struct {
	ID      string  `json:"Id"`
	Title   string  `json:"Title"`
	Desc    string  `json:"Desc"`
	Content string  `json:"Content"`
	Owner   int     `json:"Owner"`
	Viewers []uint8 `json:"Viewers"`
	Editors []uint8 `json:"Editors"`

	/*	owner INT,
		viewers INT[],
		editors INT[],*/
}
