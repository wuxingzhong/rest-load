package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/wuxingzhong/rest-parser/parser"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	var (
		restJsonConfig string
		restFile       string
	)
	app := &cli.App{
		Name:      "rest-load",
		Usage:     "rest api 负载测试",
		ArgsUsage: "rest文件",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Aliases: []string{
					"c",
				},
				Usage:       "json配置文件",
				Destination: &restJsonConfig,
				Value:       "http-client.env.json",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				restFile = c.Args().First()
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 读取配置文件
	restJsonData, err := ioutil.ReadFile(restJsonConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	varMap := make(parser.VarMap, 10)
	err = json.Unmarshal(restJsonData, &varMap)
	if err != nil {
		log.Fatal(err)
		return
	}

	restInfoList, err := parser.RestParser(restFile, varMap)
	//data, _ := json.Marshal(restInfoList)
	// fmt.Printf("%v\n\n", string(data))

	//fmt.Printf("%v\n\n", string(data))
	resultList := make(ResultList, len(restInfoList))
	for k, v := range restInfoList {
		out, err := curlCmd(resultList, &v)
		if err != nil {
			fmt.Printf("err(%v)\n", err)
		}
		resultList[k] = out
		// httpCmd(&v)
	}

	return
}

func httpCmd(restInfo *parser.RestInfo) (out string, err error) {
	args := make([]string, 0, 10)
	if len(restInfo.Body) > 0 {
		args = append(args, "--body")
		args = append(args, restInfo.Body)
	}

	args = append(args, restInfo.Method)
	args = append(args, restInfo.Path)

	if len(restInfo.Header) > 0 {
		for headerK, headerV := range restInfo.Header {
			args = append(args, fmt.Sprintf("%v:%v", headerK, headerV))
		}
	}
	fmt.Printf("%v\n", restInfo.Comment)
	out = runsCmd("http", args...)
	fmt.Printf("%v\n", out)
	return "", nil
}

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

func runsCmd(name string, arg ...string) (out string) {
	cmd := exec.Command(name, arg...)
	fmt.Printf("%v\n", cmd.String())
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	if err := cmd.Run(); err != nil {
		fmt.Printf("err(%v),%v\n", err, stderrBuf.String())
		return
	}
	out = stdoutBuf.String()
	return
}
