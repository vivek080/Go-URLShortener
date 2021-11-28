package main

import (
	"log"
	"net/http"

	"github.com/vivek080/Go-URLShortener/pkg/http/rest"
)

func main() {
	http.HandleFunc("/URLShortner", rest.ServeHTTP)
	log.Println("URL Shortner Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
