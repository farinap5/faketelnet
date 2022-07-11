package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"net/url"
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
		c.Shell(conn, &profile)
	} else if c.User != "none" && c.Pass != "none" {
		if c.User == profile.Username && c.Pass == profile.Password {
			log.Printf("profile logged ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			c.Shell(conn, &profile)
		} else {
			log.Printf("profile login error ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			conn.Write([]byte("error: access denial"))
		}
	} else if c.User != "none" {
		if c.User == profile.Username {
			log.Printf("profile logged ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			c.Shell(conn, &profile)
		} else {
			log.Printf("profile login error ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			conn.Write([]byte("error: access denial"))
		}
	} else if c.Pass != "none" {
		if c.Pass == profile.Password {
			log.Printf("profile logged ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			c.Shell(conn, &profile)
		} else {
			log.Printf("profile login error ip=%s login=%s:%s", profile.IP, profile.Username, profile.Password)
			conn.Write([]byte("error: access denial"))
		}
	}
	
	conn.Close()
	if c.Ta != "none" {
		c.taserver(profile)
	}
}


func (c Config)taserver(u User) {
	client := &http.Client{}
	req,_ := http.NewRequest("GET",c.Ta+"/api",nil)
	log.Println(c.Ta+"/api")

	cmd := b64.StdEncoding.EncodeToString([]byte(u.CMD))

	req.Header.Set("Actor-IP",strings.Split(u.IP,":")[0])
	req.Header.Set("Actor-User",url.QueryEscape(u.Username))
	req.Header.Set("Actor-Pass",url.QueryEscape(strings.TrimSuffix(u.Password,"\r")))
	req.Header.Set("Actor-CMDs",strconv.Itoa(u.NC))
	req.Header.Set("Arctor-CMD",cmd)
	req.Header.Set("Actor-Proto","Telnet")
	req.Header.Set("User-Token",c.Token)

	res,err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	} else {
		arr,_ := ioutil.ReadAll(res.Body)
		log.Println(string(arr))
	}

}