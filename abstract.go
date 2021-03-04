package doc4go

import "html/template"

type Abstract struct {
	contents []template.HTML
}

// Contents 获取内容
func (this *Abstract) Contents() []template.HTML {
	return this.contents
}

// Write 添加文本内容
func (this *Abstract) Write(lines ...string) {
	for _, line := range lines {
		this.WriteTag(P(line))
	}
}

// WriteTag 添加 HTML 标签内容
func (this *Abstract) WriteTag(tags ...Tag) {
	for _, tag := range tags {
		this.contents = append(this.contents, tag.HTML())
	}
}

// Clean 清除所有内容
func (this *Abstract) Clean() {
	this.contents = nil
}
