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
	}
	return
}
