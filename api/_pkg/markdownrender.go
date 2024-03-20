package pkg

import (
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

const (
	SEP = "---"
)

func ParseMeta(bytes []byte) (map[string]string, error) {
	var m = map[string]string{}
	var s = string(bytes)
	var f, l int

	f = 0
	l = 0

	if strings.HasPrefix(s, SEP) {
		// Read meta from bytes
		f = strings.Index(s, SEP) + len(SEP)
		l = strings.Index(s[f:], SEP) + len(SEP)
		// Split string by lines
		for _, l := range strings.Split(s[f:l], "\n") {
			if l != "" {
				// Split lines by :
				kv := strings.Split(l, ":")
				// Populate map
				var k = strings.TrimSpace(kv[0])
				var v = strings.TrimSpace(kv[1])

				m[k] = v
			}
		}
	}

	return m, nil
}

func RenderMarkdown(bytes []byte) ([]byte, error) {
	// Remove meta from bytes
	var s = string(bytes)
	if strings.HasPrefix(s, SEP) {
		// Meta to be removed
		var f = strings.LastIndex(s, SEP) + len(SEP)
		bytes = []byte(s[f:])
	}

	extensions := parser.CommonExtensions
	parser := parser.NewWithExtensions(extensions)
	doc := parser.Parse(bytes)

	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer), nil
}
