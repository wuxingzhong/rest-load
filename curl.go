package main

import (
	"fmt"
	"github.com/tidwall/sjson"
	"github.com/wuxingzhong/rest-parser/parser"
	"net/url"
	"strings"
)

func curlRun(c *Config, restInfoList []parser.RestInfo) {
	var jsonStr = `{}`
	for k, v := range restInfoList {
		extArgs := strings.Split(c.ExtArgs, " ")
		out, err := curlCmd(jsonStr, &v, extArgs)
		if err != nil {
			fmt.Printf("err(%v)\n", err)
		}
		if htype, b := v.Header["Content-Type"]; b && strings.Compare(htype, "application/json") == 0 {
			jsonStr, _ = sjson.SetRaw(jsonStr, fmt.Sprintf("list.%d.params", k), v.Body)
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
				jsonStr, _ = sjson.Set(jsonStr, fmt.Sprintf("list.%d.params", k), values2)
			}
		}
		jsonStr, _ = sjson.SetRaw(jsonStr, fmt.Sprintf("list.%d.body", k), out)
	}
}

func curlCmd(json string, restInfo *parser.RestInfo, extArgs []string) (out string, err error) {
	var (
		args []string
	)
	fmt.Printf("%d: %v\n", restInfo.Index-1, restInfo.Comment)
	args = append(args, extArgs...)
	hModel := &HttpModel{Json: json}
	if len(restInfo.Header) > 0 {
		for headerK, headerV := range restInfo.Header {
			args = append(args, "-H")
			restInfo.Header[headerK] = hModel.Replace(headerV)
			head := fmt.Sprintf("%v: %v", headerK, restInfo.Header[headerK])
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

	out = runsCmd("curl", args...)
	fmt.Printf("%v\n\n", out)
	return
}
