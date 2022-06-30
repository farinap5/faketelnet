package main

import (
	"flag"
	"fmt"
)

func main() {
	var Addr = flag.String("a", "0.0.0.0:5555", "Address and port.")
	var Bnnr = flag.String("b", "BusyBox v1.27.2", "Banner.")
	var help = flag.Bool("h", false, "Help Menu")
	var user = flag.String("u", "none", "Username")
	var pass = flag.String("p", "none", "Password")
	flag.Parse()

	if *help {
		Help()
		return
	}
	fmt.Println("#### faketelnet ####")

	c := Config{
		Addr:   *Addr,
		Banner: *Bnnr,
		User:   *user,
		Pass:   *pass,
	}

	c.Init_server()
}
