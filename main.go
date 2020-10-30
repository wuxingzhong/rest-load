package main

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
	"github.com/wuxingzhong/rest-parser/parser"
	"io/ioutil"
	"log"
	"os"
)

var (
	conf Config
)

func main() {
	app := &cli.App{
		Name:  "rest-load",
		Usage: "rest api负载测试",
		Commands: []*cli.Command{
			{
				Name:      "curl",
				Usage:     "curl方式,流程测试",
				ArgsUsage: "rest文件",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "config",
						Aliases: []string{
							"c",
						},
						Usage:       "json配置文件",
						Destination: &conf.RestJsonConfig,
						Value:       "http-client.env.json",
					},
					&cli.StringFlag{
						Name: "configEnv",
						Aliases: []string{
							"e",
						},
						Usage:       "配置环境名称",
						Destination: &conf.ConfigEnv,
						Value:       "default",
					},
					&cli.StringFlag{
						Name: "extArgs",
						Aliases: []string{
							"g",
						},
						Usage:       "外部命令额外参数",
						Destination: &conf.ExtArgs,
						Value:       "",
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						conf.RestFile = c.Args().First()
					} else {
						log.Fatalln("缺少rest文件参数")
						return nil
					}
					restInfoList := initRun(&conf)
					curlRun(&conf, restInfoList)
					return nil
				},
			},
			{
				Name:      "ab",
				Usage:     "ab方式,流程测试",
				ArgsUsage: "rest文件",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "config",
						Aliases: []string{
							"c",
						},
						Usage:       "json配置文件",
						Destination: &conf.RestJsonConfig,
						Value:       "http-client.env.json",
					},
					&cli.StringFlag{
						Name: "configEnv",
						Aliases: []string{
							"e",
						},
						Usage:       "配置环境名称",
						Destination: &conf.ConfigEnv,
						Value:       "default",
					},
					&cli.StringFlag{
						Name: "extArgs",
						Aliases: []string{
							"g",
						},
						Usage:       "外部命令额外参数",
						Destination: &conf.ExtArgs,
						Value:       "",
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						conf.RestFile = c.Args().First()
					} else {
						log.Fatalln("缺少rest文件参数")
						return nil
					}
					restInfoList := initRun(&conf)
					abRun(&conf, restInfoList)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func initRun(c *Config) []parser.RestInfo {
	// 读取配置文件
	restJsonData, err := ioutil.ReadFile(c.RestJsonConfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	nameMap := make(map[string]parser.VarMap)
	err = json.Unmarshal(restJsonData, &nameMap)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	varMap := nameMap[c.ConfigEnv]
	restInfoList, err := parser.RestParser(c.RestFile, varMap)
	//data,_ := json.MarshalIndent( restInfoList[1],"","    ",)
	//fmt.Printf("%v\n", string(data))
	return restInfoList
}
