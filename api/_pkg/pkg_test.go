package pkg

import (
	"testing"
)

func TestVercelBlobGet(t *testing.T) {
	var url string
	url = "https://pyj4ulx4cmwnqsvz.public.blob.vercel-storage.com/sample.md"
	// Initialise blob
	vb := VercelBlob{}
	b, err := vb.Get(url)

	if err != nil {
		t.Fatal(err)
	}

	if len(b) == 0 {
		t.Error("B is 0")
	}
}

func TestVercelBlobList(t *testing.T) {
	vb := VercelBlob{}
	l, err := vb.List()

	if err != nil {
		t.Fatal(err)
	}

	if len(l) == 0 {
		t.Fatal("No blobs")
	}
}
