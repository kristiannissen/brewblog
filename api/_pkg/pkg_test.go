package pkg

import (
	"testing"

	s "brewblog/_pkg/service"
)

func TestServiceNew(t *testing.T) {
	service := s.NewVercelService()

	t.Error("Yoko Oh No", service)
}

func TestServiceList(t *testing.T) {
	service := s.NewVercelService()
	list, err := service.List()

	if err != nil {
		t.Error(err)
	}

	if len(list) == 0 {
		t.Error("No blobs")
	}
}

func TestServiceFind(t *testing.T) {
	service := s.NewVercelService()
	url, err := service.Find("sample.md")

	if err != nil {
		t.Error(err)
	}

	if url == "" {
		t.Error("No URL found")
	}
}

func TestServiceDownload(t *testing.T) {
	service := s.NewVercelService()
	url, _ := service.Find("sample.md")

	if url == "" {
		t.Error("No URL Found")
	}

	bytes, err := service.Download(url)

	if err != nil {
		t.Error(err)
	}

	if len(bytes) == 0 {
		t.Error("Nothing to download")
	}
}
