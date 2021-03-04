package doc4go

import (
	"fmt"
	"html/template"
	"strings"
)

type Tag interface {
	HTML() template.HTML
}

type P string

func (this P) HTML() template.HTML {
	return template.HTML(fmt.Sprintf("<p>%s</p>", this))
}

type Pre string

func (this Pre) HTML() template.HTML {
	return template.HTML(fmt.Sprintf("<pre>%s</pre>", this))
}

type H1 string

func (this H1) HTML() template.HTML {
	return template.HTML(fmt.Sprintf("<h1>%s</h1>", this))
}

type Table struct {
	caption string
	heads   []string
	rows    [][]string
}

func NewTable() *Table {
	var t = &Table{}
	return t
}

func (this *Table) Caption(caption string) *Table {
	this.caption = caption
	return this
}

func (this *Table) Head(items ...string) *Table {
	this.heads = items

	return this
}

func (this *Table) Row(items ...string) *Table {
	this.rows = append(this.rows, items)
	return this
}

func (this *Table) HTML() template.HTML {
	var builder = &strings.Builder{}
	builder.WriteString("<table>")

	if this.caption != "" {
		builder.WriteString("<caption>")
		builder.WriteString(this.caption)
		builder.WriteString("</caption>")
	}

	builder.WriteString("<thead><tr>")
	for _, item := range this.heads {
		builder.WriteString("<th>")
		builder.WriteString(item)
		builder.WriteString("</th>")
	}
	builder.WriteString("</tr></thead>")

	builder.WriteString("<tbody>")
	for _, row := range this.rows {
		builder.WriteString("<tr>")
		for _, item := range row {
			builder.WriteString("<td>")
			builder.WriteString(item)
			builder.WriteString("</td>")
		}
		builder.WriteString("</tr>")
	}
	builder.WriteString("</tbody>")
	builder.WriteString("</table>")
	return template.HTML(builder.String())
}

func (this *Table) Clean() {
	this.rows = nil
}
