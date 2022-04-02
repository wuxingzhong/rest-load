package main

import (
	"fmt"
	"text/template"
	"time"
)

// 时间戳
func Timestamp() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}

func FuncMap() template.FuncMap {
	t := template.FuncMap{
		"timestamp": Timestamp,
	}
	return t
}
