package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func handleSession(listener net.Listener) error {
	//
	// Accept session.
	//

	conn, err := listener.Accept()
	if err != nil {
		return err
	}

	defer conn.Close()

	//
	// Handle session messages.
	//

	for {
		//
		// Read message.
		//

		buffer := make([]byte, 1024)
		if _, err := conn.Read(buffer); err != nil {
			if err == io.EOF {
				fmt.Println("eof")
				return nil
			}

			return err
		}

		fmt.Printf("[server] %s\n", buffer)

		//
		// Respond.
		//

		if _, err := conn.Write([]byte("pong")); err != nil {
			return err
		}
	}
}

func serveTCP() error {
	//
	// Start listening.
	//

	listener, err := net.ListenTCP("tcp",
		&net.TCPAddr{
			IP:   net.ParseIP("localhost"),
			Port: port,
		},
	)
	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		if err := handleSession(listener); err != nil {
			return err
		}

		fmt.Println("[server] session done")
		fmt.Println()
	}
}

func createSessionAndDoFinitePings() error {
	conn, err := net.DialTCP("tcp",
		nil,
		&net.TCPAddr{
			IP:   net.ParseIP("localhost"),
			Port: port,
		},
	)
	if err != nil {
		return err
	}

	defer conn.Close()

	for i := 0; i < 5; i++ {
		if _, err := conn.Write([]byte("ping")); err != nil {
			return err
		}

		buffer := make([]byte, 1024)
		if _, err := conn.Read(buffer); err != nil {
			return err
		}

		fmt.Printf("[client] %s\n", buffer)

		time.Sleep(1 * time.Second)
	}

	return nil
}

func doPingsTCP() error {
	for {
		createSessionAndDoFinitePings()
	}
}
