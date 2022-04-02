package main

type Config struct {
	// json 配置文件
	RestJsonConfig string
	// 配置文件环境
	ConfigEnv string
	// rest文件
	RestFile string
	// 外部命令额外参数
	ExtArgs string
}
