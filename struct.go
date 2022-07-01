package main

type Config struct {
	Addr    string
	Banner  string
	User    string
	Pass    string
	BusyBox string
}

type User struct {
	Username string
	Password string
	IP       string
}
