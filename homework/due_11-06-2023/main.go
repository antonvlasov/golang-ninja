package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	flushInterval = time.Duration(1) * time.Second
	maxQueueSize  = 1000000
	UDPPacketSize = 1500
	port          = 4001
	alphanum      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	secondsToRun  = 60
)

var address string
var flushTicker *time.Ticker
var nbWorkers int

// var loading = true
var seconds uint64 = 0

func init() {
	flag.StringVar(&address, "addr", fmt.Sprintf(":%d", port), "Address of the UDP server to test")
	flag.IntVar(&nbWorkers, "concurrency", runtime.NumCPU(), "Number of workers to run in parallel")
}

type message struct {
	addr   net.Addr
	msg    []byte
	length int
}

func generateMessage(n int) []byte {
	bytes := make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}

func main() {
	fmt.Println("----------start-------")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			//loading = false
			runtime.Gosched()
			atomic.AddUint64(&totalWriter, opsWriter)
			atomic.AddUint64(&totalReader, opsReader)
			log.Println("----------finish-----------")
			log.Printf("Total messages sent %d", totalWriter)
			log.Printf("Total messages received %d", totalReader)
			percentage := float64(totalWriter-totalReader) / float64(totalWriter)
			log.Printf("Packet loss percentage: %%%f", percentage)
			os.Exit(0)
		}
	}()

	go runServer()
	go runClient()

	flushTicker = time.NewTicker(flushInterval)
	for range flushTicker.C {
		log.Printf("[info]seconds: %d \n", seconds)
		atomic.AddUint64(&seconds, 1)
		if seconds > secondsToRun {
			c <- os.Interrupt
		}
	}

	select {}
}

//
