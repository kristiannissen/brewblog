package service

import "brewblog/_pkg/domain"

type GoogleService struct{}

func NewGoogleService() BlobService {
	return &GoogleService{}
}

func (v *GoogleService) Download(url string) ([]byte, error) {
	return []byte(``), nil
}

func (v *GoogleService) List() ([]domain.Blob, error) {
	return []domain.Blob{}, nil
}

func (v *GoogleService) Find(pathname string) (string, error) {
	return "", nil
}
