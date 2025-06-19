package shortener

import (
	"fmt"
	"sync"
)

type SingleThreadedShortener struct {
	urls map[string]string
	mu   sync.RWMutex
}

// validate we've implemented the UrlShortener interface
var _ UrlShortener = (*SingleThreadedShortener)(nil)

func NewSingleThreadedShortener() UrlShortener {
	return &SingleThreadedShortener{
		urls: make(map[string]string),
	}
}

func (s *SingleThreadedShortener) Shorten(longURL string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return generateAndAdd(s.urls, longURL)
}

func (s *SingleThreadedShortener) Expand(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if dest, found := s.urls[shortURL]; found {
		return dest, nil
	}
	return "", fmt.Errorf("short URL not found")
}
