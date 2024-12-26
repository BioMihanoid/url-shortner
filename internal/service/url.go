package service

import "github.com/BioMihanoid/url-shortner/internal/repository"

type urlService struct {
	repo repository.URL
}

func newUrlService(repo repository.URL) *urlService {
	return &urlService{repo: repo}
}

func (s *urlService) SaveUrl(sUrl, lUrl string) error {
	return s.repo.SaveUrl(sUrl, lUrl)
}

func (s *urlService) GetLongUrl(sUrl string) (string, error) {
	return s.repo.GetLongUrl(sUrl)
}
