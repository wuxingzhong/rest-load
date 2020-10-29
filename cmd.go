package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

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
