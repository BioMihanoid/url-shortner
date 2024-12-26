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

func (u *UrlPostgres) SaveUrl(sUrl, lUrl string) error {
	query := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2)", urlTable, shortUrl, longUrl)

	_, err := u.db.Exec(query, sUrl, lUrl)

	return err
}

func (u *UrlPostgres) GetLongUrl(sUrl string) (string, error) {
	var lUrl string

	query := fmt.Sprintf("SELECT %s FROM %s WHERE short_url=$1", longUrl, urlTable)

	err := u.db.Get(&lUrl, query, sUrl)

	return lUrl, err
}
