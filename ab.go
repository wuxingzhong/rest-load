package main

import (
	"fmt"
	"github.com/wuxingzhong/rest-parser/parser"
	"io/ioutil"
	"strings"
)

func abRun(c *Config, restInfoList []parser.RestInfo) {
	json := `{}`
	for _, v := range restInfoList {
		extArgs := strings.Split(c.ExtArgs, " ")
		_, err := abCmd(json, &v, extArgs)
		if err != nil {
			fmt.Printf("err(%v)\n", err)
		}
		// 处理out
	}
}

func abCmd(json string, restInfo *parser.RestInfo, extArgs []string) (out string, err error) {
	var (
		args []string
	)
	args = append(args, extArgs...)
	hModel := &HttpModel{Json: json}
	if len(restInfo.Header) > 0 {
		for headerK, headerV := range restInfo.Header {
			args = append(args, "-H")
			head := hModel.Replace(fmt.Sprintf("%v: %v", headerK, headerV))
			args = append(args, head)
		}
	}
	if len(restInfo.Body) > 0 {
		args = append(args, "-p")
		body := hModel.Replace(restInfo.Body)
		_ = ioutil.WriteFile("/tmp/test.data", []byte(body), 0666)
		args = append(args, "/tmp/test.data")
	}
	path := hModel.Replace(restInfo.Path)
	args = append(args, path)
	fmt.Printf("%d: %v %v\n", restInfo.Index, restInfo.Method, restInfo.Comment)

	out = runsCmd("ab", args...)
	fmt.Printf("\n%v\n", out)
	return
}
