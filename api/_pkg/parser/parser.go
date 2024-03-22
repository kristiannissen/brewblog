package parser

import (
	"errors"
	"log"
	"strings"

	"brewblog/_pkg/domain"
)

const (
	SEP = "---"
	LB  = "\n"
	DLB = "\n\n"
)

func StripMeta(b []byte) []byte {
	s := string(b)

	if strings.HasPrefix(s, SEP) {
		// Meta found
		s = s[strings.LastIndex(s, SEP)+len(SEP):]
	}

	return []byte(strings.TrimSpace(s))
}

// FIXME: Not happy about the naming...
func ParseJSON(b []byte) (domain.Article, error) {
	var article domain.Article
	var meta map[string]string

	// Parse meta
	meta = ParseMeta(b)
	article.Meta = meta
	// Strip meta
	b = StripMeta(b)
	var s = string(b)

	if strings.Index(s, "#") < 0 {
		log.Println("No title")
		return article, errors.New("No title in document")
	}
	// ParseHeader
	article.Title = ParseTitle(b)
	// Update s
	s = s[strings.Index(s, article.Title)+len(article.Title):]
	// Split s by DLB
	for _, p := range strings.Split(s, DLB) {
		if len(p) > 0 {
			para := domain.Paragraph{}
			// Check if there is a header
			if strings.HasPrefix(p, "#") {
				para.Header = p[:strings.Index(s, LB)]
			} else if strings.HasPrefix(p, "!") {
				//
			} else {
				// Plain text
				para.Body = strings.TrimSpace(p)
			}
			// Ensure no empty paras
			article.Paragraphs = append(article.Paragraphs, para)
		}
	}

	return article, nil
}

func ParseMeta(b []byte) map[string]string {
	// M will contain all meta key/value pairs
	var m = map[string]string{}
	var s = string(b)
	var f, l int

	f = 0
	l = 0

	if strings.HasPrefix(s, SEP) {
		// Read meta from string
		f = strings.Index(s, SEP) + len(SEP)
		l = strings.Index(s[f:], SEP) + len(SEP)

		for _, line := range strings.Split(s[f:l], "\n") {
			// Sprint string by line
			if line != "" {
				// Split lne by kv
				kv := strings.Split(line, ":")
				// Populate map
				m[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
			}
		}
	}

	return m
}

func ParseTitle(b []byte) string {
	var t, s string

	s = string(b)

	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, "#") {
			t = line
			break
		}
	}
	// Remove #
	t = t[strings.Index(t, "#")+len("#"):]

	return strings.TrimSpace(t)
}
