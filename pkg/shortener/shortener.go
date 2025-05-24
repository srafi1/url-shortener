package shortener

type UrlShortener interface {
	Shorten(url string) (string, error)
	Expand(shortenedUrl string) (string, error)
}
