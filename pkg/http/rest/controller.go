package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	httpJson "github.com/vivek080/Go-URLShortener/pkg/http/rest/json"
	"github.com/vivek080/Go-URLShortener/pkg/storage"
)

// ServeHTTP servers the request based on the method called
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createURLShortener(w, r)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// createURLShortener creates a short URL for the given url.
func createURLShortener(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request httpJson.URLShortenerData
	var response httpJson.URLShortenerResponseData

	w.Header().Add("X-Request-ID", r.Header.Get("X-Request-ID"))
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		httpJson.SendHTTPErrorResponse(w, http.StatusUnprocessableEntity, "Unable to read the request", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		log.Print(err.Error())
		httpJson.SendHTTPErrorResponse(w, http.StatusBadRequest, "Unable to unmarshal the request", err.Error())
		return
	}

	if request.Data.URL == "" {
		log.Print("URL is empty")
		httpJson.SendHTTPErrorResponse(w, http.StatusBadRequest, "URL must not be empty", "URL is empty")
		return
	}

	shortURL, err := storage.CreateURLShortener(request.Data.URL)
	if err != nil {
		log.Print("Unable to creates a short URL", "err", err)
		httpJson.SendHTTPErrorResponse(w, http.StatusInternalServerError, "Unable to creates a short URL", err.Error())
		return
	}

	response.MapFromURLShortener(shortURL)

	httpJson.SendHTTPSuccessResponse(w, http.StatusCreated, response)
}
