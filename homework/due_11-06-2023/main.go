package main

import (
	"crypto/rand"
)

const (
	port = 4001
)

func generateMessage() []byte {
	res := make([]byte, 64)
	_, _ = rand.Read(res)
	return res
}

func main() {
	go serveUDP()
	go doPingsUDP()

	select {}
}
