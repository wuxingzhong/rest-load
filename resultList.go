package main

import (
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"strings"
)

type ResultList []string

func (r ResultList) Replace(old string) string {
	reg := regexp.MustCompile("\\${(.+?)}")
	str := reg.ReplaceAllStringFunc(old, r.replaceFunc)
	return str
}

func (r ResultList) replaceFunc(old string) string {
	old = strings.ReplaceAll(old, "${", "")
	old = strings.ReplaceAll(old, "}", "")
	tmpStr := strings.SplitN(old, ".", 2)
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
