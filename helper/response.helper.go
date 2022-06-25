package helper

import (
	"encoding/json"
	"net/http"
)

type helper struct {
}

type Helper interface {
	ResponseHandler(w http.ResponseWriter, data interface{}, message string, httpstatus int)
}

//NewCategoryRepository create new instance for categry repository
func NewHelper() Helper {
	return &helper{}
}

//ResponseHandler generate response object
func (h *helper) ResponseHandler(w http.ResponseWriter, data interface{}, message string, httpstatus int) {

	response := make(map[string]interface{})
	sucess := false
	if httpstatus == 200 {
		sucess = true
	}
	response = map[string]interface{}{"status": sucess, "message": message}
	response["data"] = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpstatus)

	json.NewEncoder(w).Encode(response)
}
