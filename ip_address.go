package main

import (
	"net"
)

type IPAddress struct {
	ipAddress net.IPNet
	description string
}