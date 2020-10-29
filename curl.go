package main

import (
	"fmt"
	"github.com/wuxingzhong/rest-parser/parser"
	"strings"
)

func curlRun(c *Config, restInfoList []parser.RestInfo) {
	resultList := make(ResultList, len(restInfoList))
	for k, v := range restInfoList {
		extArgs := strings.Split(c.ExtArgs, " ")
		out, err := curlCmd(resultList, &v, extArgs)
		if err != nil {
			fmt.Printf("err(%v)\n", err)
		}
		resultList[k] = out
	}
}

func curlCmd(resList ResultList, restInfo *parser.RestInfo, extArgs []string) (out string, err error) {
	var (
		args []string
	)
	args = append(args, extArgs...)
	if len(restInfo.Header) > 0 {
		for headerK, headerV := range restInfo.Header {
			args = append(args, "-H")
			head := resList.Replace(fmt.Sprintf("%v: %v", headerK, headerV))
			args = append(args, head)
		}
	}
	if len(restInfo.Body) > 0 {
		args = append(args, "-d")
		body := resList.Replace(restInfo.Body)
		args = append(args, body)
	}
	args = append(args, "-X")
	args = append(args, restInfo.Method)
	path := resList.Replace(restInfo.Path)
	args = append(args, path)
	fmt.Printf("%d: %v %v\n", restInfo.Index, restInfo.Method, restInfo.Comment)

	out = runsCmd("curl", args...)
	fmt.Printf("\n%v\n", out)
	return
}
