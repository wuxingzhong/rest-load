package main

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"text/template"
)

func Template(src string) string {

	tpl := template.Must(template.New("base").Delims("{%", "%}").Funcs(sprig.TxtFuncMap()).Parse(src))

	var buffer bytes.Buffer
	_ = tpl.Execute(&buffer, "")
	return buffer.String()
}
