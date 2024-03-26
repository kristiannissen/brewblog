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
<body data-page="post">
<header class="container_header header_logo" next-page-hide>
<h1 class="font_heading"><a href="/">Hansen<span class="font_amp">&</span>Nissen</a></h1>
</header>
<main class="container__content">
	<article class="container__article column">
		<div class="content__header content__navigation pl-1 pr-1">
			<nav class="navigation__breadcrumb">
				<ul class="breadcrumb__list">
					<li class="breadcrumb__list_item"><a href="/">Hjem</a></li>
					<li class="breadcrumb__list_item">Blog Post</li>
				</ul>
			</nav>
			<h1 class="font_header mb-2">{{title}}</h1>
		</div>
					{{.}}
	</template>
</article>
</main>
<footer class="container__footer">
	<div class="column">
		<b>Stærke øl til svage mænd!</b>
	</div>
	<div class="column">2</div>
</footer>
</body>
</html>
`

	var buff bytes.Buffer
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		return []byte(buff.String()), err
	}
	err = t.Execute(&buff, template.HTML(string(b)))
	if err != nil {
		return []byte(buff.String()), err
	}

	return []byte(strings.TrimSpace(buff.String())), nil
}
