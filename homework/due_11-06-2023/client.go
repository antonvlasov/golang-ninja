package main

import (
	"flag"
	"log"
	mrand "math/rand"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var bufferPoolWriter sync.Pool
var opsWriter uint64 = 0
var totalWriter uint64 = 0
var flushTickerWriter *time.Ticker

func runClient() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	bufferPoolWriter = sync.Pool{
		New: func() interface{} { return make([]byte, UDPPacketSize) },
	}
	load(nbWorkers)

	flushTickerWriter = time.NewTicker(flushInterval)
	for range flushTickerWriter.C {
		log.Printf("[client]Ops/s %f", float64(opsWriter)/flushInterval.Seconds())
		atomic.AddUint64(&totalWriter, opsWriter)
		atomic.StoreUint64(&opsWriter, 0)
	}
}

func load(maxWorkers int) error {
	delay := 500
	for i := 0; i < maxWorkers; i++ {
		go func() {
			for seconds < secondsToRun {
				mrand.Seed(time.Now().Unix())
				n := mrand.Intn(UDPPacketSize - 1)
				write(generateMessage(n), n)
				time.Sleep(time.Duration(delay) * time.Microsecond)
			}
		}()
	}
	return nil
}

func write(buf []byte, n int) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Printf("Error connecting to server: %s", err)
		return
	}
	defer conn.Close()
	defer func() { bufferPoolWriter.Put(buf) }()

	_, err = conn.Write(buf[0:n])
	if err != nil {
		log.Printf("Error sending to server: %s", err)
		return
	}
	atomic.AddUint64(&opsWriter, 1)
}
