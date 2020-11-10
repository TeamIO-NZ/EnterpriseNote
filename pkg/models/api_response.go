package models

//APIResponse struct for api response
type APIResponse struct {
	Code    int         `json:"code,omitempty"`    //the response code
	Message string      `json:"message,omitempty"` //any message you want to put
	Data    interface{} `json:"data,omitempty"`    //data to send to the thing
}

//BuildAPIResponseSuccess function for building all those api responses im going to write
func BuildAPIResponseSuccess(message string, data interface{}) APIResponse {
	var api APIResponse
	api.Code = 200
	api.Message = message
	api.Data = data
	return api
}

//BuildAPIResponseFail function for building all those api responses im going to write
func BuildAPIResponseFail(message string, data interface{}) APIResponse {
	var api APIResponse
	api.Code = 400
	api.Message = message
	api.Data = data
	return api
}

//BuildAPIResponseUnauthorized function for building all those api responses im going to write
func BuildAPIResponseUnauthorized(message string) APIResponse {
	var api APIResponse
	api.Code = 401
	api.Message = message
	return api
}
