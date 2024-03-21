package render

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func RenderMarkdown(b []byte) []byte {

	ext := parser.CommonExtensions
	parser := parser.NewWithExtensions(ext)
	doc := parser.Parse(b)

	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func RenderTemplate(b []byte) ([]byte, error) {
	var tpl = `
<!DOCTYPE html>
<html>
<body>{{.}}</body>
</html>
`

	var buff bytes.Buffer
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		return []byte(buff.String()), err
	}
	err = t.Execute(&buff, string(b))
	if err != nil {
		return []byte(buff.String()), err
	}

	return []byte(strings.TrimSpace(buff.String())), nil
}
