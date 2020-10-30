package main

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"text/template"
)

type ResultList []string

func (r ResultList) Replace(old string) string {
	tpl := template.Must(template.New("base").Delims("{%", "%}").
		Funcs(sprig.TxtFuncMap()).
		Funcs(map[string]interface{}{
			"result": r.Result,
		}).
		Funcs(FuncMap()).Parse(old))
	var buffer bytes.Buffer
	_ = tpl.Execute(&buffer, r)
	return buffer.String()
}

func (r ResultList) Result(expStr string) string {

	tmpStr := strings.SplitN(expStr, ".", 2)
	v := ""
	if len(tmpStr) < 2 {
		return ""
	}
	if index, err := strconv.Atoi(tmpStr[0]); err != nil {
		return ""
	} else if index < len(r) {

		value := gjson.Get(r[index-1], tmpStr[1])
		v = value.String()
	}
	return v
}
