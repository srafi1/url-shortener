package main

import (
	"log"
	"maps"
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

var shortenerImpls = map[string]func() shortener.UrlShortener{
	"single-threaded": shortener.NewSingleThreadedShortener,
	"single-producer": shortener.NewSingleProducerShortener,
}

func main() {
	log.Println("URL shortener is running...")
	log.Printf("Listening at: %s", LISTEN_ADDR)

	var s shortener.UrlShortener
	if len(os.Args) > 1 {
		if newShortenerFn, ok := shortenerImpls[os.Args[1]]; ok {
			s = newShortenerFn()
		} else {
			keys := make([]string, 0)
			for k := range maps.Keys(shortenerImpls) {
				keys = append(keys, k)
			}
			log.Printf("shortener arg not recognized: %s\nvalid options are: %v", os.Args[1], keys)
			return
		}
	} else {
		s = shortener.NewSingleThreadedShortener()
	}

	router := routing.GetRouter(s)
	if err := http.ListenAndServe(LISTEN_ADDR, router); err != nil {
		log.Printf("failed to serve http: %s", err.Error())
	}
}
