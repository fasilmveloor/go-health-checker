package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[Down] %v is unreachable, \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[Up] %v is reachable, \n From: %v \n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
