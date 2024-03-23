package parser

import (
	"errors"
	"log"
	"regexp"
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
		if strings.TrimSpace(p) != "" {
			para := domain.Paragraph{}
			// Check if there is a header
			if strings.HasPrefix(p, "#") {
				para.Header = p[:strings.Index(p, LB)]
				// log.Println("Header", para.Header)
			} else if strings.HasPrefix(p, "!") {
				// Check if image has alt and title
				if m, err := regexp.MatchString(`^!\[(.*)?\]\((.*)?\s"(.*)"?\)`, p); err == nil {
					var image domain.Image
					// Find images
					if m == true {
						// Alt + title
						pat := regexp.MustCompile(`^!\[(.*)?\]\((.*)?\s"(.*)?"\)`)
						submatch := pat.FindStringSubmatch(p)
						// Populate struct
						image.Title = submatch[1]
						image.URL = submatch[2]
						image.Alt = submatch[3]
					} else {
						// Alt - title
						pat := regexp.MustCompile(`^!\[(.*)?\]\((.*)?\)`)
						submatch := pat.FindStringSubmatch(p)
						image.Title = submatch[1]
						image.URL = submatch[2]
					}
					para.Images = append(para.Images, image)
				}
			} else {
				// Plain text
				para.Body = strings.TrimSpace(p)
				log.Println(para.Body)
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
