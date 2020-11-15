package web

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"net/http"      // used to access the request and response object of the api

	_ "github.com/lib/pq" //sql driver. blank is required
	"go.iosoftworks.com/EnterpriseNote/pkg/models"
)

//------------------------------JSON Webrequests Hander functions -- Specifics --------------------------------//

//CheckConnection Checks the database is connected
func (server Server) CheckConnection(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.BuildAPIResponseSuccess("Success", nil)) //Send a 200 response so we know its connected
}
