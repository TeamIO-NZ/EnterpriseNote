package tests

import (
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

var Users = []models.User{
	models.User{
		Name:     "Lithial",
		Password: "1234",
		Email:    "me@james.me",
	},
	models.User{
		Name:     "Joe",
		Password: "1234",
		Email:    "you@james.me",
	},
	models.User{
		Name:     "Peter",
		Password: "1234",
		Email:    "us@james.me",
	},
	models.User{
		Name:     "Arran",
		Password: "1234",
		Email:    "re@james.me",
	},
	models.User{
		Name:     "Finn",
		Password: "1234",
		Email:    "de@james.me",
	},
	models.User{
		Name:     "Sam",
		Password: "1234",
		Email:    "la@james.me",
	},
	models.User{
		Name:     "Sam",
		Password: "1234",
		Email:    "ke@james.me",
	},
}

var Notes = []models.Note{
	models.Note{
		ID:      "0",
		Title:   "James is the overlord",
		Desc:    "The best overlord",
		Content: "The very best overlord there is",
		Owner:   1,
		Viewers: []int{1, 2, 3},
		Editors: []int{4, 5},
	},
	models.Note{
		ID:      "1",
		Title:   "Joe is the Minion",
		Desc:    "The best minion",
		Content: "So i decree",
		Owner:   1,
		Viewers: []int{1, 2, 3},
		Editors: []int{4, 5},
	},
	models.Note{
		ID:      "2",
		Title:   "No joe is the boss",
		Desc:    "The best boss",
		Content: "So i decree",
		Owner:   2,
		Viewers: []int{6, 2, 3},
		Editors: []int{4, 5},
	},
}

func testGetNote() {

}
func testGetAllNotes() {

}
func testUpdateNote() {

}
func testDeleteNote() {

}
func testInsertNote() {

}
func testGetUser() {

}
func testGetUserByName() {

}
func testGetAllUsers() {

}
func testUpdateUser() {

}
func testDeleteUser() {

}
func testInsertUser() {

}
func testGetSpecificNotes() {

}
func testLogin() {

}
func testGetAllNotesUserHasAccessTo() {

}
