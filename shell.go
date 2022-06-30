package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

// Shell pseudo shell
func (c Config) Shell(conn net.Conn, p User) {
	var shell string
	if p.Username == "root" {
		shell = "~# "
	} else {
		shell = "~$ "
	}
	conn.Write([]byte(shell))

	var cmd string
	var err error
	for {
		cmd, err = bufio.NewReader(conn).ReadString('\r')
		if err != nil {
			log.Println(err.Error())
			break
		}
		cmd = strings.TrimSuffix(cmd, "\r")
		log.Printf("input ip=%s login=%s cmd=%s", p.IP, p.Username, cmd)
		if cmd == "exit" || cmd == "quit" || cmd == "C^" {
			conn.Close()
			break
		} else {
			r := p.proc(cmd)
			conn.Write([]byte(r + "\n" + shell))
		}
	}
}

func (u User) proc(c string) string {
	var ret string
	cs := strings.Split(c, " ")

	if c == "whoami" {
		ret = u.Username
	} else if c == "pwd" {
		ret = "/home/" + c
	} else if c == "ls" {
		ret = "README.md\nscript.py"
	} else if cs[0] == "echo" && len(cs) >= 2 {
		ret = cs[1]
	} else {
		ret = "sh: 1: " + c + ": not found"
	}

	return ret
}
