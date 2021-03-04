package doc4go

import (
	"html/template"
	"os"
)

type Document struct {
	title    string     // 标题
	tpl      Template   // 模板
	abstract *Abstract  // 摘要
	chapters []*Chapter // 章节
}

func New(title string, tpl Template) *Document {
	var d = &Document{}
	d.title = title
	if tpl == nil {
		tpl = NewDefaultTemplate()
	}
	d.tpl = tpl
	return d
}

// Title 获取标题
func (this *Document) Title() string {
	return this.title
}

// Abstracts 获取摘要
func (this *Document) Abstract() *Abstract {
	if this.abstract == nil {
		this.abstract = &Abstract{}
	}
	return this.abstract
}

// Chapters 获取章节
func (this *Document) Chapters() []*Chapter {
	return this.chapters
}

// Chapter 根据标题获取章节
// 如果章节不存在，则会根据提供的标题创建新的章节
func (this *Document) Chapter(title string) *Chapter {
	for _, chapter := range this.chapters {
		if chapter.title == title {
			return chapter
		}
	}

	var nChapter = &Chapter{}
	nChapter.title = title
	this.chapters = append(this.chapters, nChapter)
	return nChapter
}

// WriteToFile 保存到文件
func (this *Document) WriteToFile(file string) error {

	var tpl = template.New(this.title)
	tpl = template.Must(tpl.Parse(this.tpl.Content()))

	var nFile, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_SYNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer nFile.Close()
	nFile.Truncate(0)

	if err = tpl.Execute(nFile, this); err != nil {
		return err
	}

	return nil
}

// Clean 清除所有内容
func (this *Document) Clean() {
	this.abstract = nil
	this.chapters = nil
}
