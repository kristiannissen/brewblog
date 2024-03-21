package service

import "time"

// TODO: Move to seperate folder
type Blob struct {
	URL         string
	PathName    string
	Size        uint64
	UploadedAt  time.Time
	ContentType string
}

type BlobService interface {
	// Get
	Download(url string) ([]byte, error)
	// List
	List() ([]Blob, error)
	// Find returens URL
	Find(patname string) (string, error)
}
