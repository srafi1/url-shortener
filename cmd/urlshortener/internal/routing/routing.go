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

func writeJSON(w http.ResponseWriter, status int, resp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func writeText(w http.ResponseWriter, status int, resp string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	if _, err := w.Write([]byte(resp)); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func ServeHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("404: unmatched route for path=%s", r.URL.String())
		writeJSON(w, http.StatusNotFound, Response{Error: "resource not found"})
		return
	}

	writeText(w, http.StatusOK, "hello world")
}

func ServeShorten(s shortener.UrlShortener) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		originalUrl := r.URL.Path[len("/shorten/"):]
		if originalUrl == "" {
			writeJSON(w, http.StatusBadRequest, Response{Error: "url not found in request"})
			return
		}

		shortenedUrl, err := s.Shorten(originalUrl)
		if err != nil {
			log.Printf("failed to shorten url: %s", err)
			writeJSON(w, http.StatusInternalServerError, Response{Error: "failed to shorten url"})
			return
		}

		writeJSON(w, http.StatusOK, Response{
			ShortenedUrl: shortenedUrl,
		})
	}
}

func ServeExpand(s shortener.UrlShortener) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		shortenedUrl := r.URL.Path[len("/expand/"):]
		if shortenedUrl == "" {
			writeJSON(w, http.StatusBadRequest, Response{Error: "url not found in request"})
			return
		}

		originalUrl, err := s.Expand(shortenedUrl)
		if err != nil {
			log.Printf("failed to expand url: %s", err)
			writeJSON(w, http.StatusInternalServerError, Response{Error: "failed to expand url"})
			return
		}

		writeJSON(w, http.StatusOK, Response{
			OriginalUrl: originalUrl,
		})
	}
}
