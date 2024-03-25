package service

import (
	"log"

	"brewblog/_pkg/domain"
	s "brewblog/_pkg/service"

	"github.com/rpdg/vercel_blob"
)

type VercelService struct {
	client *vercel_blob.VercelBlobClient
}

func (v *VercelService) NewService() s.BlobService {
	return &VercelService{
		client: vercel_blob.NewVercelBlobClient(),
	}
}

func (v *VercelService) Download(url string) ([]byte, error) {
	b, err := v.client.Download(url, vercel_blob.DownloadCommandOptions{})
	if err != nil {
		log.Println(err)
		return b, err
	}
	return b, nil
}

func (v *VercelService) Find(pathname string) (string, error) {
	blobs, err := v.List()
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Find matching pathname
	var url string
	for _, blob := range blobs {
		if blob.PathName == pathname {
			url = blob.URL
			break
		}
	}

	return url, nil
}

func (v *VercelService) List() ([]domain.Blob, error) {
	files, err := v.client.List(vercel_blob.ListCommandOptions{})

	if err != nil {
		// Log error
		log.Println(err)
		return []domain.Blob{}, err
	}

	var blobs []domain.Blob
	for _, f := range files.Blobs {
		blobs = append(blobs, domain.Blob{
			URL:        f.URL,
			PathName:   f.PathName,
			Size:       f.Size,
			UploadedAt: f.UploadedAt,
		})
	}

	return blobs, nil
}
