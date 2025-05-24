package routing

import (
	"log"
	"net/http"
)

func GetRouter() http.Handler {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", serveHello)
	return mux
}

func serveHello(w http.ResponseWriter, _ *http.Request) {
	response := "hello world"
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("failed to write hello response: %w", err)
	}
}
