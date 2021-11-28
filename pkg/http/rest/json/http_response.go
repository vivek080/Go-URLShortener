package json

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

// SendHTTPSuccessResponse sends a HTTP response with reponse body.
func SendHTTPSuccessResponse(w http.ResponseWriter, status int, response interface{}) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Print("Unable to Marshal reponse body to HTTP Response", "error", err.Error())
	}
}

// SendHTTPErrorResponse sends a HTTP response with a single error.
func SendHTTPErrorResponse(w http.ResponseWriter, status int, code string, description string) {

	var jerr ErrorData
	jerr.Data.Status = strconv.Itoa(status)
	jerr.Data.Code = code
	jerr.Data.Description = description

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(jerr)
	if err != nil {
		log.Print("Unable to Marshal error body to HTTP Response", "error", err.Error())
	}
}
