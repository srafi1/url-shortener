package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/srafi1/url-shortener/cmd/urlshortener/internal/routing"
)

var LISTEN_ADDR = func() string {
	if addr, ok := os.LookupEnv("LISTEN_ADDR"); ok {
		return addr
	}
	return "0.0.0.0:3000"
}()

func main() {
	log.Println("URL shortener is running...")
	log.Printf("Listening at: %s", LISTEN_ADDR)

	s := &shortenerImpl{}
	router := routing.GetRouter(s)
	if err := http.ListenAndServe(LISTEN_ADDR, router); err != nil {
		log.Printf("failed to serve http: %s", err.Error())
	}
}

type shortenerImpl struct{}

// Expand implements shortener.UrlShortener.
func (s *shortenerImpl) Expand(shortenedUrl string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

// Shorten implements shortener.UrlShortener.
func (s *shortenerImpl) Shorten(url string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}
