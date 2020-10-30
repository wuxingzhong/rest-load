package main

import (
	"fmt"
	"github.com/tidwall/sjson"
	"github.com/wuxingzhong/rest-parser/parser"
	"net/url"
	"strings"
)

func curlRun(c *Config, restInfoList []parser.RestInfo) {
	json := `{}`
	for k, v := range restInfoList {
		extArgs := strings.Split(c.ExtArgs, " ")
		out, err := curlCmd(json, &v, extArgs)
		if err != nil {
			fmt.Printf("err(%v)\n", err)
		}
		if htype, b := v.Header["Content-Type"]; b && strings.Compare(htype, "application/json") == 0 {
			json, _ = sjson.SetRaw(json, fmt.Sprintf("list.%d.params", k), v.Body)
		} else {
			rawQuery := ""
			if strings.Compare(strings.ToLower(v.Method), "get") == 0 {
				if u, err := url.ParseRequestURI(v.Path); err == nil {
					rawQuery = u.RawQuery
				}
			} else {
				rawQuery = v.Body
			}
			if values, err := url.ParseQuery(rawQuery); err == nil {
				values2 := make(map[string]string, len(values))
				for k1, v1 := range values {
					values2[k1] = strings.Join(v1, ",")
				}
				json, _ = sjson.Set(json, fmt.Sprintf("list.%d.params", k), values2)
			}
		}
		json, _ = sjson.SetRaw(json, fmt.Sprintf("list.%d.body", k), out)
	}
}

func curlCmd(json string, restInfo *parser.RestInfo, extArgs []string) (out string, err error) {
	var (
		args []string
	)
	args = append(args, extArgs...)
	hModel := &HttpModel{Json: json}
	if len(restInfo.Header) > 0 {
		for headerK, headerV := range restInfo.Header {
			args = append(args, "-H")
			restInfo.Header[headerK] = hModel.Replace(headerV)
			head := fmt.Sprintf("%v: %v", headerK, headerV)
			args = append(args, head)
		}
	}
	if len(restInfo.Body) > 0 {
		args = append(args, "-d")
		restInfo.Body = hModel.Replace(restInfo.Body)
		args = append(args, restInfo.Body)
	}
	args = append(args, "-X")
	args = append(args, restInfo.Method)
	restInfo.Path = hModel.Replace(restInfo.Path)
	args = append(args, restInfo.Path)
	fmt.Printf("%d: %v %v\n", restInfo.Index, restInfo.Method, restInfo.Comment)

	out = runsCmd("curl", args...)
	fmt.Printf("\n%v\n", out)
	return
}
