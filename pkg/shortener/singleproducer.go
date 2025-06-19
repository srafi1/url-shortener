package shortener

import (
	"fmt"
	"sync"
)

type SingleProducerShortener struct {
	urls     map[string]string
	requests chan *request
}

func NewSingleProducerShortener() UrlShortener {
	s := &SingleProducerShortener{
		urls:     map[string]string{},
		requests: make(chan *request),
	}
	go s.startProducer()
	return s
}

type request struct {
	s  slot
	wg sync.WaitGroup
}

type slot struct {
	key, val string
	err      error
}

func (s *SingleProducerShortener) startProducer() {
	for r := range s.requests {
		s.handleRequest(r)
	}
}

func (s *SingleProducerShortener) handleRequest(r *request) {
	defer r.wg.Done()

	// NOTE: a better check for saturation is the count of active URLs vs the generateFriendlyID probability space
	for retries := 5; retries > 0; retries -= 1 {
		short := generateFriendlyID()
		if _, found := s.urls[short]; !found {
			s.urls[short] = r.s.key
			r.s.val = short
			return
		}
	}

	r.s.err = fmt.Errorf("memory is too saturated")
	return
}

// validate we've implemented the UrlShortener interface
var _ UrlShortener = (*SingleProducerShortener)(nil)

// Shorten implements UrlShortener.
func (s *SingleProducerShortener) Shorten(url string) (string, error) {
	r := &request{
		s: slot{
			key: url,
		},
		wg: sync.WaitGroup{},
	}
	r.wg.Add(1)
	s.requests <- r
	r.wg.Wait()
	return r.s.val, r.s.err
}

// Expand implements UrlShortener.
func (s *SingleProducerShortener) Expand(shortenedUrl string) (string, error) {
	if dest, found := s.urls[shortenedUrl]; found {
		return dest, nil
	}
	return "", fmt.Errorf("short URL not found")
}
