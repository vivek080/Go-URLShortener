package json

// URLShortenerData represents request to create a new URLShortner
type URLShortenerData struct {
	Data URLShortenerRequest `json:"data"`
}

// URLShortenerRequest represents a URL
type URLShortenerRequest struct {
	URL string `json:"url" validate:"required"`
}
