package service

import "brewblog/_pkg/domain"

type GoogleService struct{}

func (g *GoogleService) NewService() BlobService {
	return &GoogleService{}
}

func (g *GoogleService) Download(url string) ([]byte, error) {
	return []byte(``), nil
}

func (g *GoogleService) List() ([]domain.Blob, error) {
	return []domain.Blob{}, nil
}

func (g *GoogleService) Find(pathname string) (string, error) {
	return "", nil
}
