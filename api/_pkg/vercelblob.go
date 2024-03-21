package pkg

import (
	"log"

	v "github.com/rpdg/vercel_blob"
)

/*
 * Implements interface from storage
 */
type VercelBlob struct {
	// TODO: Add client
}

func (b *VercelBlob) Find(pathname string) (string, error) {
	l, _ := b.List()
	for _, k := range l {
		if k.PathName == pathname {
			return k.URL, nil
		}
	}

	return "", nil
}

/*
 * Downloads bytes
 */
func (b *VercelBlob) Get(url string) ([]byte, error) {
	// Initialise client
	client := v.NewVercelBlobClient()
	// Try to download bytes
	bytes, err := client.Download(url, v.DownloadCommandOptions{})
	// Error handling
	if err != nil {
		log.Println("VercelBlob", err)
		return []byte(``), err
	}
	// Return data
	return bytes, nil
}

func (b *VercelBlob) List() ([]Blob, error) {
	// Initialize client
	client := v.NewVercelBlobClient()
	files, err := client.List(v.ListCommandOptions{})

	if err != nil {
		log.Println("VercelBlob", err)
		return []Blob{}, err
	}

	blobs := []Blob{}
	for _, f := range files.Blobs {
		blobs = append(blobs, Blob{
			URL:        f.URL,
			PathName:   f.PathName,
			Size:       f.Size,
			UploadedAt: f.UploadedAt,
		})
	}

	return blobs, nil
}
