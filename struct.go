package main

type Config struct {
	Addr    string
	Banner  string
	User    string
	Pass    string
	BusyBox string
	Ta		string //ThreatActor Server
	Token	string
}

type User struct {
	Username string
	Password string
	IP       string
	CMD		 string
	NC 		 int // number of commands runned
}
