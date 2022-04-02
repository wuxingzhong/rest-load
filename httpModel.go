package main

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/tidwall/gjson"
	"strings"
	"text/template"
)

type HttpModel struct {
	Json string
}

func (h *HttpModel) Replace(old string) string {
	if !strings.Contains(old, "{%") {
		return old
	}
	// fmt.Printf("----%v\n", old)
	tpl := template.Must(template.New("base").Delims("{%", "%}").
		Funcs(sprig.TxtFuncMap()).
		Funcs(map[string]interface{}{
			"http": h.Http,
		}).
		Funcs(FuncMap()).Parse(old))
	var buffer bytes.Buffer
	_ = tpl.Execute(&buffer, h)
	return buffer.String()
}

func (h *HttpModel) Http(expStr string) string {

	expStr = fmt.Sprintf("list.%v", expStr)
	value := gjson.Get(h.Json, expStr)
	return value.String()
}
