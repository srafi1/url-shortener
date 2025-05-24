package shortener

import (
	"fmt"
)

type SingleThreadedShortener struct {
	urls map[string]string
}

// validate we've implemented the UrlShortener interface
var _ UrlShortener = (*SingleThreadedShortener)(nil)

func NewSingleThreadedShortener() *SingleThreadedShortener {
	return &SingleThreadedShortener{
		urls: make(map[string]string),
	}
}

func (s *SingleThreadedShortener) Shorten(longURL string) (string, error) {
	// TODO: implement me
	return longURL, nil
}

func (s *SingleThreadedShortener) Expand(shortURL string) (string, error) {
	if dest, found := s.urls[shortURL]; found {
		return dest, nil
	}
	return "", fmt.Errorf("short URL not found")
}
