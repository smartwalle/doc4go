package doc4go

type Chapter struct {
	title    string     // 标题
	abstract *Abstract  // 摘要
	sections []*Section // 小节
}

// Title 获取标题
func (this *Chapter) Title() string {
	return this.title
}

// Abstracts 获取摘要
func (this *Chapter) Abstract() *Abstract {
	if this.abstract == nil {
		this.abstract = &Abstract{}
	}
	return this.abstract
}

// Sections 获取小节
func (this *Chapter) Sections() []*Section {
	return this.sections
}

// Section 根据标题获取小节
// 如果小节不存在，则会根据提供的标题创建新的小节
func (this *Chapter) Section(title string) *Section {
	for _, section := range this.sections {
		if section.title == title {
			return section
		}
	}

	var nSection = &Section{}
	nSection.title = title
	this.sections = append(this.sections, nSection)
	return nSection
}

// Clean 清除所有内容
func (this *Chapter) Clean() {
	this.abstract = nil
	this.sections = nil
}
