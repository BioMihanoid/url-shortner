package service

import "github.com/BioMihanoid/url-shortner/internal/repository"

type URL interface {
	SaveUrl(sUrl, lUrls string) error
	GetLongUrl(sUrl string) (string, error)
}

type Service struct {
	URL
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		URL: newUrlService(repos.URL),
	}
}
