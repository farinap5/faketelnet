package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func (c Config) Init_server() {
	var id int = 1
	serve, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	println("Listenig", c.Addr)
	defer serve.Close()
	for {
		client, err := serve.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go c.handle_conn(client)
		id++
	}
}

func (c Config) handle_conn(conn net.Conn) {
	profile := User{
		IP: conn.RemoteAddr().String(),
	}
	var err error

	conn.Write([]byte(c.Banner + "\nstg login: "))
	profile.Username, err = bufio.NewReader(conn).ReadString('\r')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	profile.Username = strings.TrimSuffix(profile.Username, "\r")

	conn.Write([]byte("password: "))
	profile.Password, err = bufio.NewReader(conn).ReadString('\r')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	if c.User == "none" && c.Pass == "none" {
		log.Printf("profile logged ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
		c.Shell(conn, profile)
	} else if c.User != "none" && c.Pass != "none" {
		if c.User == profile.Username && c.Pass == profile.Password {
			log.Printf("profile logged ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			c.Shell(conn, profile)
		} else {
			log.Printf("profile login error ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			conn.Write([]byte("error: access denial"))
		}
	} else if c.User != "none" {
		if c.User == profile.Username {
			log.Printf("profile logged ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			c.Shell(conn, profile)
		} else {
			log.Printf("profile login error ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			conn.Write([]byte("error: access denial"))
		}
	} else if c.Pass != "none" {
		if c.Pass == profile.Password {
			log.Printf("profile logged ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			c.Shell(conn, profile)
		} else {
			log.Printf("profile login error ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			conn.Write([]byte("error: access denial"))
		}
	}

	conn.Close()
}
