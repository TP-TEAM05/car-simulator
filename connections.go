package main

import (
	"fmt"
	"net"
)

type Connection struct {
	Id            string
	UDPConnection *net.UDPConn
	OtherAddress  *net.UDPAddr
	NextSendIndex int
}

func (connection *Connection) WriteDatagram(data []byte) {
	_, _ = connection.UDPConnection.Write(data)
	fmt.Printf("Sending message to %v: %s\n", connection.OtherAddress, data[:96])
	connection.NextSendIndex++
}
