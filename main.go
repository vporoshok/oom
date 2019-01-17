package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	speed = flag.String("speed", "100M", "Size add per second. Unit: B, K, M, G")
)

func main() {
	flag.Parse()
	var (
		k int
		u byte
	)
	_, err := fmt.Sscanf(*speed, "%d%c", &k, &u)
	if err != nil {
		log.Fatalf("invalid speed: %s", err)
	}
	if k <= 0 {
		log.Fatal("invalid spped: must be positive")
	}
	switch u {
	case 'K':
		k *= 1024
	case 'M':
		k *= 1024 * 1024
	case 'G':
		k *= 1024 * 1024 * 1024
	}

	var buffers [][]byte

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	t := time.NewTicker(time.Second)

	for {
		select {
		case <-c:
			log.Print("interrupted")
			os.Exit(0)

		case <-t.C:
			log.Printf("increase buffer: %d", len(buffers)*k)
			buffers = append(buffers, make([]byte, k))
		}
	}
}
