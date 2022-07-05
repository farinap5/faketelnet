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
	var busy = flag.String("busy", "none", "Busybox path.")
	var tact = flag.String("ta","none","ThreatActors Server.")
	var tken = flag.String("t","none","Token")
	flag.Parse()

	if *help {
		Help()
		return
	}
	fmt.Println("#### faketelnet ####")

	c := Config{
		Addr:    *Addr,
		Banner:  *Bnnr,
		User:    *user,
		Pass:    *pass,
		BusyBox: *busy,
		Ta:      *tact,
		Token:   *tken,
	}

	c.Init_server()
}
