package main

import (
	"github.com/smartwalle/doc4go"
)

func main() {
	var doc = doc4go.New("Mir 游戏引擎脚本指令说明文档", nil)
	doc.Abstract().Write("本文档用于说明 Mir 游戏引擎中支持的脚本指令及各指令的使用方法.")
	doc.Abstract().Write("后续内容纯属扯淡。")
	doc.Abstract().Write("The Go programming language is an open source project to make programmers more productive.")
	doc.Abstract().Write("Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.")

	var chapter = doc.Chapter("用户")
	chapter.Abstract().Write("用户相关的接口")

	var section = chapter.Section("GET - /api/v1/user/:user_id")
	section.Write("获取用户信息")
	var tbl = doc4go.NewTable()
	tbl.Caption("参数列表")
	tbl.Head("名称", "类型", "是否必须", "默认值", "描述")
	tbl.Row("user_id", "int", "是", "无", "用户的 id")
	section.WriteTag(tbl)

	section = chapter.Section("POST - /api/v1/user/refresh_token")
	section.Write("更新当前用户 Token")
	tbl.Clean()
	tbl.Row("refresh_token", "string", "是", "无", "用户当前的 refresh_token")
	section.WriteTag(tbl)

	chapter = doc.Chapter("指令")
	section = chapter.Section("PLAYER_VAR_SET")
	section.Write("设置玩家变量的值")
	section.WriteTag(doc4go.Pre(`
#ACT
PLAYER_VAR_SET var_name value
`))

	section = chapter.Section("$PLAYER_VAR_GET")
	section.Write("获取玩家变量的值")
	section.WriteTag(doc4go.Pre(`
#SAY
变量的值为: <$PLAYER_VAR_GET|var_name>
`))

	doc.Chapter("指令").
		Section("PLAYER_VAR_CHECK").
		Write("比较玩家变量的值, 示例：PLAYER_VAR_CHECK var_name > 10").
		Write("asdfasfasdf").
		Write("asdfasfasdf").
		Write("asdfasfasdf")

	doc.WriteToFile("doc.html")
}
