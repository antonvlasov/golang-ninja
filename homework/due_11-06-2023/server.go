package main

import (
	"flag"
	"log"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type messageQueue chan message

var mq messageQueue
var bufferPoolReader sync.Pool
var opsReader uint64 = 0
var totalReader uint64 = 0
var flushTickerReader *time.Ticker

func (mq messageQueue) enqueue(m message) {
	mq <- m
}

func (mq messageQueue) dequeue() {
	for m := range mq {
		handleMessage(m.addr, m.msg[0:m.length])
		bufferPoolReader.Put(m.msg)
	}
}

func runServer() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	bufferPoolReader = sync.Pool{
		New: func() interface{} { return make([]byte, UDPPacketSize) },
	}
	mq = make(messageQueue, maxQueueSize)
	listenAndReceive(nbWorkers)

	flushTickerReader = time.NewTicker(flushInterval)
	for range flushTickerReader.C {
		log.Printf("[server]Ops/s %f", float64(opsReader)/flushInterval.Seconds())
		atomic.AddUint64(&totalReader, opsReader)
		atomic.StoreUint64(&opsReader, 0)
	}
}

func listenAndReceive(maxWorkers int) error {
	c, err := net.ListenPacket("udp", address)
	if err != nil {
		return err
	}
	for i := 0; i < maxWorkers; i++ {
		go mq.dequeue()
		go receive(c)
	}
	return nil
}

// receive accepts incoming datagrams on c and calls handleMessage() for each message
func receive(c net.PacketConn) {
	defer c.Close()

	for {
		msg := bufferPoolReader.Get().([]byte)
		nbytes, addr, err := c.ReadFrom(msg[0:])
		if err != nil {
			log.Printf("Error %s", err)
			continue
		}
		mq.enqueue(message{addr, msg, nbytes})
	}
}

func handleMessage(addr net.Addr, msg []byte) {
	//fmt.Println(string(msg))
	atomic.AddUint64(&opsReader, 1)
}
