package json

import (
	"github.com/vivek080/Go-URLShortener/pkg/storage"
)

// ErrorData represents Error object
type ErrorData struct {
	Data Error `json:"data"`
}

// Error is object for application errors response
type Error struct {
	Code        string `json:"errorCode"`
	Description string `json:"errorDescription,omitempty"`
	Status      string `json:"httpStatus"`
}

// URLShortenerResponseData represents response of the URLShortener
type URLShortenerResponseData struct {
	Data URLShortenerResponse `json:"data"`
}

// URLShortenerResponse represents a URL shortener struct
type URLShortenerResponse struct {
	URL      string `json:"url"`
	ShortURL string `json:"shortURL"`
}

// MapFromURLShortener maps fields from urlShortener to response object
func (response *URLShortenerResponseData) MapFromURLShortener(url *storage.URLShortener) {
	response.Data.URL = url.URL
	response.Data.ShortURL = url.ShortURL

}
