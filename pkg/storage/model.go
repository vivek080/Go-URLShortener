package storage

// URLShortener represents a URL shortener struct
type URLShortener struct {
	URL      string `json:"url"`
	ShortURL string `json:"shortURL"`
}
