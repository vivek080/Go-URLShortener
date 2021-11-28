package rest

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	httpJson "github.com/vivek080/Go-URLShortener/pkg/http/rest/json"
)

// TestCreateShortURL_1 test tries creating a short URL & expects status code be 201
func TestCreateShortURL_1(t *testing.T) {
	os.Setenv("URL_SHORTENER_FILE_PATH", "../../../urlShortener.txt")
	os.Setenv("BASE_URL", "infra.go/")
	w := httptest.NewRecorder()
	jsonBody := `{
		"data":{
			"url":"test.com"
		}
	}`
	r := httptest.NewRequest("POST", "/URLShortner", strings.NewReader(jsonBody))

	createURLShortener(w, r)
	assert.Equal(t, 201, w.Result().StatusCode)
}

// TestCreateShortURLError_1 test fails creating a short URL due to read body error & expects status code be 422
func TestCreateShortURLError_1(t *testing.T) {
	os.Setenv("URL_SHORTENER_FILE_PATH", "../../../urlShortener.txt")
	os.Setenv("BASE_URL", "infra.go/")
	w := httptest.NewRecorder()

	r := httptest.NewRequest("POST", "/URLShortner", httpJson.ErrReader(0))

	createURLShortener(w, r)
	assert.Equal(t, 422, w.Result().StatusCode)
}

// TestCreateShortURLError_2 test fails creating a short URL due to json body error & expects status code be 400
func TestCreateShortURLError_2(t *testing.T) {
	w := httptest.NewRecorder()
	jsonBody := `{
		"data":{
			"url":"test.com"
		
	}`
	r := httptest.NewRequest("POST", "/URLShortner", strings.NewReader(jsonBody))

	createURLShortener(w, r)
	assert.Equal(t, 400, w.Result().StatusCode)
}

// TestCreateShortURLError_3 test fails creating a short URL due to missing field value & expects status code be 400
func TestCreateShortURLError_3(t *testing.T) {
	w := httptest.NewRecorder()
	jsonBody := `{
		"data":{
			"url":""
		}
	}`
	r := httptest.NewRequest("POST", "/URLShortner", strings.NewReader(jsonBody))

	createURLShortener(w, r)
	assert.Equal(t, 400, w.Result().StatusCode)
}

// TestCreateShortURLError_4 test fails creating a short URL due to CreateURLShortener failure & expects status code be 500
func TestCreateShortURLError_4(t *testing.T) {
	os.Setenv("URL_SHORTENER_FILE_PATH", "../../../urlShortener.tx")
	os.Setenv("BASE_URL", "infra.go/")
	w := httptest.NewRecorder()
	jsonBody := `{
		"data":{
			"url":"test.com"
		}
	}`
	r := httptest.NewRequest("POST", "/URLShortner", strings.NewReader(jsonBody))

	createURLShortener(w, r)
	assert.Equal(t, 500, w.Result().StatusCode)
}
