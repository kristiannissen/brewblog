package entry

import (
	"os"
	"path/filepath"
	"testing"
)

func setup() {
	// log.Println("ENV", os.Getenv("BLOB_READ_WRITE_TOKEN"))
}

func TestMain(m *testing.M) {
	setup()
	c := m.Run()
	os.Exit(c)
}

func TestGetEntry(t *testing.T) {
	_, err := GetEntry("sample")

	if err != nil {
		t.Fatal(err)
	}
}

func TestParseEntryData(t *testing.T) {
	var err error
	var b []byte
	var wd string
	var e Entry

	wd, err = os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	p := filepath.Join(wd, "/../../_content/sample.md")

	b, err = os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}

	e, err = ParseEntryData(string(b))

	// Test meta
	// log.Println("Meta:", e.Meta)
	if e.Meta[0].Key != "tags" {
		t.Fatal("No tags")
	}

	// Test headline
	// log.Println("Title:", e.Title)
	if e.Title == "" {
		t.Fatal("No title")
	}

	// Paragraphs
	// log.Println("Paras:", e.Paragraphs)
	if len(e.Paragraphs) == 0 {
		t.Fatal("No paragraphs")
	}
}
