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
		// Skip empty lines
		if m, err := regexp.MatchString(`^\S`, p); err == nil {
			if m == true {
				// No white space
				var para domain.Paragraph

				switch string(p[0]) {
				case "!":
					log.Println(parseImage(p))
				case "#":
					para.Header = parseHeader(p)
				default:
					para.Body = strings.TrimSpace(p)
				}
				// Populate article paragraphs
				article.Paragraphs = append(article.Paragraphs, para)
			}
		} else {
			log.Println(err)
		}
	}

	return article, nil
}

func parseImage(s string) map[string]string {
	images := map[string]string{}
	// Split by LB
	// ![The San Juan Mountains are beautiful!](/assets/images/san-juan-mountains.jpg)
	for _, l := range strings.Split(s, LB) {
		var f, e int
		// Parse title
		f = strings.Index(l, "[") + 1
		e = strings.LastIndex(l, "]")
		images["title"] = l[f:e]
		// URL
		f = strings.Index(l[e:], "(") + 1

		images["url"] = l[f:]

	}

	return images
}

func parseHeader(s string) string {
	return s
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
