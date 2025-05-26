package routing

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/srafi1/url-shortener/pkg/shortener"
)

type handler func(http.ResponseWriter, *http.Request)

func GetRouter(s shortener.UrlShortener) http.Handler {
	mux := &http.ServeMux{}
	mux.HandleFunc("/shorten/", ServeShorten(s)) // handles /shorten/<long>
	mux.HandleFunc("/expand/", ServeExpand(s))   // handles /expand/<short>
	mux.HandleFunc("/", ServeHello)              // fallback catch-all

	return mux
}

type Response struct {
	OriginalUrl  string `json:"originalUrl,omitempty"`
	ShortenedUrl string `json:"shortenedUrl,omitempty"`
	Error        string `json:"error,omitempty"`
}

func ServeHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("404: unmatched route for path=%s", r.URL.String())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := Response{Error: "resource not found"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("failed to write response: %v", err)
		}

		return
	}

	response := "hello world"
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("failed to write hello response: %s", err.Error())
	}
}

func ServeShorten(s shortener.UrlShortener) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		originalUrl := r.URL.Path[len("/shorten/"):]
		if originalUrl == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Error: "url not found in request"})
			return
		}

		shortenedUrl, err := s.Shorten(originalUrl)
		if err != nil {
			log.Printf("failed to shorten url: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Error: "failed to shorten url"})
			return
		}

		json.NewEncoder(w).Encode(Response{
			ShortenedUrl: shortenedUrl,
		})
	}
}

func ServeExpand(s shortener.UrlShortener) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		shortenedUrl := r.URL.Path[len("/expand/"):]
		if shortenedUrl == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Error: "url not found in request"})
			return
		}

		originalUrl, err := s.Expand(shortenedUrl)
		if err != nil {
			log.Printf("failed to expand url: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Error: "failed to shorten url"})
			return
		}

		json.NewEncoder(w).Encode(Response{
			OriginalUrl: originalUrl,
		})
	}
}
