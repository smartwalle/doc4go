package doc4go

import "html/template"

type Section struct {
	title    string          // 标题
	contents []template.HTML // 内容
}

// Title 获取标题
func (this *Section) Title() string {
	return this.title
}

// Contents 获取内容
func (this *Section) Contents() []template.HTML {
	return this.contents
}

// Write 添加文本内容
func (this *Section) Write(lines ...string) *Section {
	for _, line := range lines {
		this.WriteTag(P(line))
	}
	return this
}

// WriteTag 添加 HTML 标签内容
func (this *Section) WriteTag(tags ...Tag) *Section {
	for _, tag := range tags {
		this.contents = append(this.contents, tag.HTML())
	}
	return this
}

// Clean 清除所有内容
func (this *Section) Clean() {
	this.contents = nil
}
