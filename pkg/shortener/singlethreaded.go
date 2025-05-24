package shortener

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
	// TODO: implement me
	return shortURL, nil
}
