package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	userName := os.Args[1]

	args := []string{"-u", userName}

	cmd := exec.Command("/usr/bin/sacct", args...)
	cmdOut, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(cmdOut))
	}
}
