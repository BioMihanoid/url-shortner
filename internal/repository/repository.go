package repository

import "github.com/jmoiron/sqlx"

const (
	urlTable = "urls"
	longUrl  = "long_url"
)

type URL interface {
	SaveUrl(shortUrl, longUrls string) error
	GetLongUrl(shortUrl string) (string, error)
}

type Repository struct {
	URL
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		URL: NewUrlPostgres(db),
	}
}
