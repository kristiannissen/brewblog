package service

// Implementation of service interface
/*
	// Get
	Get(pathname string) ([]byte, error)
	// List
	List() ([]Blob, error)
	// Find
	Find(patname string) (bool, error)
*/

type GoogleService struct{}

func NewGoogleService() BlobService {
	return &GoogleService{}
}

func (v *GoogleService) Download(url string) ([]byte, error) {
	return []byte(``), nil
}

func (v *GoogleService) List() ([]Blob, error) {
	return []Blob{}, nil
}

func (v *GoogleService) Find(pathname string) (string, error) {
	return "", nil
}
