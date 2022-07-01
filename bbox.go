package main

import (
	"fmt"
	"os/exec"
)

func (conf Config) Bbox(c string) string {
	cmd := exec.Command(conf.BusyBox, c)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return c + ": applet not found"
	}
	return string(stdout)
}
