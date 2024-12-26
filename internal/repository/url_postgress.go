package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UrlPostgres struct {
	db *sqlx.DB
}

func NewUrlPostgres(db *sqlx.DB) *UrlPostgres {
	return &UrlPostgres{db: db}
}

func (u *UrlPostgres) SaveUrl(shortUrl, longUrls string) error {
	query := fmt.Sprintf("INSERT INTO %s (short_url, long_url) VALUES ($1, $2)", urlTable)

	_, err := u.db.Exec(query, shortUrl, longUrls)

	return err
}

func (u *UrlPostgres) GetLongUrl(shortUrl string) (string, error) {
	var lUrl string

	query := fmt.Sprintf("SELECT %s FROM %s WHERE short_url=$1", longUrl, urlTable)

	err := u.db.Get(&lUrl, query, shortUrl)

	return lUrl, err
}
