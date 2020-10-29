package main

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"text/template"
	"time"
)

type MockTemplate struct {
}

// 时间戳
func Timestamp() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}

func FuncMap() template.FuncMap {
	t := template.FuncMap{
		"timestamp": Timestamp,
	}
	return t
}

func (m *MockTemplate) Template(src string) string {

	tpl := template.Must(template.New("base").Delims("{%", "%}").Funcs(sprig.TxtFuncMap()).Funcs(FuncMap()).Parse(src))

	var buffer bytes.Buffer
	_ = tpl.Execute(&buffer, m)
	return buffer.String()
}
