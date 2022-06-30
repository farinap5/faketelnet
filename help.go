package main

import (
	"fmt"
	"github.com/cheynewallace/tabby"
)

func Help() {
	fmt.Println("#### faketelnet ####")
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION", "DEFAULT")
	t.AddLine("-h", "Help menu.", "False")
	t.AddLine("-b", "Specify banner.", "BusyBox v1.27.2")
	t.AddLine("-a", "Address Host:port.", "0.0.0.0:5555")
	t.AddLine("-u", "User for login.", "none")
	t.AddLine("-p", "Password for login.", "none")
	print("\n")
	t.Print()
	print("\n")
}
