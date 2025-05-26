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
	mux.HandleFunc("/", ServeHello)
	mux.HandleFunc("/shorten", ServeShorten(s))
	mux.HandleFunc("/expand", ServeExpand(s))
	return mux
}

type Response struct {
	OriginalUrl  string `json:"originalUrl,omitempty"`
	ShortenedUrl string `json:"shortenedUrl,omitempty"`
	Error        string `json:"error,omitempty"`
}

func ServeHello(w http.ResponseWriter, _ *http.Request) {
	response := "hello world"
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("failed to write hello response: %s", err.Error())
	}
}

func ServeShorten(s shortener.UrlShortener) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		originalUrl := r.URL.Query().Get("url")
		if originalUrl == "" {
			w.WriteHeader(http.StatusBadRequest)
			response, _ := json.Marshal(Response{Error: "url not found in request"})
			if _, err := w.Write(response); err != nil {
				log.Printf("failed to write error: %s", err.Error())
			}
			return
		}

		shortenedUrl, err := s.Shorten(originalUrl)
		if err != nil {
			log.Printf("failed to shorten url: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			response, _ := json.Marshal(Response{Error: "failed to shorten url"})
			if _, werr := w.Write(response); werr != nil {
				log.Printf("failed to write error: %s", werr.Error())
			}
			return
		}

		response, _ := json.Marshal(Response{OriginalUrl: originalUrl, ShortenedUrl: shortenedUrl})
		if _, err := w.Write(response); err != nil {
			log.Printf("failed to write shorten response: %s", err.Error())
		}
	}
}

func ServeExpand(s shortener.UrlShortener) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		shortenedUrl := r.URL.Query().Get("url")
		if shortenedUrl == "" {
			w.WriteHeader(http.StatusBadRequest)
			response, _ := json.Marshal(Response{Error: "url not found in request"})
			if _, err := w.Write(response); err != nil {
				log.Printf("failed to write error: %s", err.Error())
			}
			return
		}

		originalUrl, err := s.Expand(shortenedUrl)
		if err != nil {
			log.Printf("failed to expand url: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			response, _ := json.Marshal(Response{Error: "failed to expand url"})
			if _, werr := w.Write(response); werr != nil {
				log.Printf("failed to write error: %s", werr.Error())
			}
			return
		}

		response, _ := json.Marshal(Response{OriginalUrl: originalUrl, ShortenedUrl: shortenedUrl})
		if _, err := w.Write(response); err != nil {
			log.Printf("failed to write expand response: %s", err.Error())
		}
	}
}
