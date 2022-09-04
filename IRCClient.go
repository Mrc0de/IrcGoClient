package main

import "net"

type Client struct {
	IrcServerConns []*net.Conn
	ClientProfiles []ClientProfile
}

// ClientProfile is the settings for each associated IRC connection
type ClientProfile struct {
	Nickname     string
	EnableIdentD bool
	WhoisConfig  WhoisConfiguration
}

type WhoisConfiguration struct {
}
