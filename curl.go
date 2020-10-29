package main

import (
	"fmt"
	"github.com/wuxingzhong/rest-parser/parser"
)

func curlCmd(resList ResultList, restInfo *parser.RestInfo) (out string, err error) {
	args := []string{
		"-k",
		"-v",
	}
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
