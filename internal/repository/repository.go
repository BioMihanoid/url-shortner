package repository

import "github.com/jmoiron/sqlx"

const (
	urlTable = "urls"
	longUrl  = "long_url"
	shortUrl = "sort_url"
)

type URL interface {
	SaveUrl(sUrl, lUrls string) error
	GetLongUrl(sUrl string) (string, error)
}

type Repository struct {
	URL
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		URL: NewUrlPostgres(db),
	}
}
