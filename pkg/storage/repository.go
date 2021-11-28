package storage

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/itchyny/base58-go"
)

// CreateURLShortener checks if the shortURL exist in the file or not
// if it exist then returns the value or else creates a new shortURL and save it in the file
func CreateURLShortener(url string) (*URLShortener, error) {
	var urlShortener URLShortener
	urlShortener.URL = url

	url = strings.TrimSpace(url)
	baseURL := os.Getenv("BASE_URL")

	urlFile := os.Getenv("URL_SHORTENER_FILE_PATH")
	file, err := os.OpenFile(urlFile, os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// checks for the url if it exist in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ":|:")
		dbURL := strings.TrimSpace(items[0])
		if dbURL == url {
			urlShortener.ShortURL = baseURL + strings.TrimSpace(items[1])
			return &urlShortener, nil
		}
	}

	shortURL := GenerateShortLink(url)
	// shortURL = uuid.New().String()[:8]

	// write the url and shortURL to the file
	fileURL := url + ":|:" + shortURL
	_, err = file.WriteString(fileURL + "\n")
	if err != nil {
		return nil, err
	}
	urlShortener.ShortURL = baseURL + shortURL

	return &urlShortener, nil
}

// GenerateShortLink generates random string for the given url.
func GenerateShortLink(initialLink string) string {
	urlHashBytes := sha256Of(initialLink)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
