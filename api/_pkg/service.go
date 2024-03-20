package pkg

import (
	"bytes"
	"log"
	"text/template"
)

/*
 * Blob Methods
 */
type BlobService interface {
	Get(url string) ([]byte, error)
	List() ([]Blob, error)
}

/*
 * Render Methods
 */
type RenderService interface {
	ParseMeta(bytes []byte) (map[string]string, error)
	RenderMarkdown(bytes []byte) ([]byte, error)
}

func RenderHTML(b []byte) (string, error) {
	var tpl = `
<!DOCTYPE html>
<html>
	<head>
		<title></title>
	</head>
	<body>{{.}}</body>
</html>
`

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Println("HTMLRender", err)
		return "", err
	}

	buff := bytes.Buffer{}
	_ = t.Execute(&buff, string(b))

	return buff.String(), nil
}
