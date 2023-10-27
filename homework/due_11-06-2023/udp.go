package main

import (
	"fmt"
	"net"
	"time"
)

func serveUDP() error {
	//
	// Start listening.
	//

	conn, err := net.ListenUDP("udp",
		&net.UDPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: port,
		},
	)
	if err != nil {
		return err
	}

	defer conn.Close()

	for {
		//
		// Read message.
		//

		buffer := make([]byte, 1024)
		_, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			return err
		}

		fmt.Printf("[server] %s\n", buffer)

		//
		// Respond.
		//

		if _, err := conn.WriteToUDP([]byte("pong"), addr); err != nil {
			return err
		}
	}
}

func doPingsUDP() error {
	conn, err := net.DialUDP("udp",
		nil,
		&net.UDPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: port,
		},
	)
	if err != nil {
		return err
	}

	defer conn.Close()

	for {
		if _, err := conn.Write([]byte("ping")); err != nil {
			return err
		}

		buffer := make([]byte, 1024)
		if _, err := conn.Read(buffer); err != nil {
			return err
		}

		fmt.Printf("[client] %s\n", buffer)

		time.Sleep(500 * time.Millisecond)
	}
}
