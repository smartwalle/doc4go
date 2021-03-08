package doc4go

type Template interface {
	Content() string
}

type DefaultTemplate struct {
	content string
}

func NewDefaultTemplate() Template {
	var t = &DefaultTemplate{}
	t.content = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>{{.Title}}</title>
  <style>
      * {
          font-family: sans-serif;
      }

      h1, h2, h3, h4 {
          margin: 1.25rem 0 1.25rem;
          color: #202224;
          font-weight: 600;
          overflow-wrap: break-word;
      }

      h1 {
          font-size: 1.75rem;
          line-height: 1;
      }

      h2 {
          font-size: 1.25rem;
          line-height: 1.25;
          background: #e0ebf5;
          padding: 0.5rem;
      }

      h3, h4 {
          margin: 1.25rem 0.35rem;
      }

      h3 {
          font-size: 1.25rem;
          line-height: 1.25;
          color: #007d9c;
      }

      p {
          margin: 0.75rem 1.25rem;
          /*max-width: 50rem;*/
          word-wrap: break-word;
          display: block;
      }

      pre {
          background: #efefef;
          padding: 0.625rem;
          border-radius: 0.3125rem;
          margin: 1.25rem;
          line-height: 1.4;
          overflow-x: auto;
          font-family: Menlo, monospace;
          font-size: 0.875rem;
      }


      table {
          border-collapse: collapse;
          margin: 1.25rem;
          text-align: center;
          border: none;
          border-spacing: 2px;
      }

      table td, table th {
          color: #666;
          font-size: 1rem;
          line-height: 1.5;
          padding: 5px 20px;
      }

      table thead th {
          background-color: #3f3f3f;
          color: #fff;
      }

      tbody tr {
          border-bottom: 1px solid #ccc;
      }

      table tr:nth-child(odd) {
          background: #fff;
      }

      table tr:nth-child(even) {
          background: #f2f2f2;
      }

      table caption {
          text-align: left;
          padding: 0.25rem 0;
          font-size: 1rem;
          line-height: 1.5;
      }

      .container {
          text-align: left;
          margin-left: auto;
          margin-right: auto;
          padding: 0 1.25rem;
      }

      .chapter {
      }

      .section {
      }

  </style>
</head>
<body>
<main id="page">
  <div class="container">
    <h1>{{.Title}}</h1>
    {{ range .Abstract.Contents }}
    {{ . }}
    {{ end }}

    {{ range .Chapters }}
    <div class="chapter">
      <h2>{{ .Title }}</h2>
      {{ range .Abstract.Contents }}
      {{ . }}
      {{ end }}
      {{ range .Sections }}
      <div class="section">
        <h3>{{ .Title }}</h3>
        {{ range .Contents }}
        {{ . }}
        {{ end }}
      </div>
      {{ end }}
    </div>
    {{ end }}
  </div>
</main>
</body>
</html>
`
	return t
}

func (this *DefaultTemplate) Content() string {
	return this.content
}
