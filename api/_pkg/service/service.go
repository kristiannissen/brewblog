package service

import "brewblog/_pkg/domain"

type BlobService interface {
	// Get
	Download(url string) ([]byte, error)
	// List
	List() ([]domain.Blob, error)
	// Find returens URL
	Find(patname string) (string, error)
}
