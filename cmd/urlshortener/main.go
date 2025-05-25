package main

import (
	"log"
	"net/http"
	"os"

	"github.com/srafi1/url-shortener/cmd/urlshortener/internal/routing"
	"github.com/srafi1/url-shortener/pkg/shortener"
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

	s := shortener.NewSingleThreadedShortener()
	router := routing.GetRouter(s)
	if err := http.ListenAndServe(LISTEN_ADDR, router); err != nil {
		log.Printf("failed to serve http: %s", err.Error())
	}
}
