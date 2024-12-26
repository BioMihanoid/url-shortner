package models

type URL struct {
	ID       int    `json:"id"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}
