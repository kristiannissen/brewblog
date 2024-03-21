package parser

import (
	"strings"
)

const (
	SEP = "---"
)

func StripMeta(b []byte) []byte {
	s := string(b)

	if strings.HasPrefix(s, SEP) {
		// Meta found
		s = s[strings.LastIndex(s, SEP)+len(SEP):]
	}

	return []byte(strings.TrimSpace(s))
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
