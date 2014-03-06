package main

import (
	"fmt"
	"net"
	"flag"
	"io/ioutil"
	"time"
)

var (
	server *string
	file *string
	rate *int
	nuance *int
)

func main() {
	server = flag.String("server", "localhost:4443", "Server to flood.")
	file = flag.String("file", "/dev/stdin", "Content to send to server.")
	rate = flag.Int("rate", 0, "Max bytes per second to attempt.")
	nuance = flag.Int("nuance", 180000, "Nuance.")

	flag.Parse()

	bytes, err := ioutil.ReadFile(*file)

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTimeout("udp", *server, 1e9)

	if err != nil {
		fmt.Println("Error establishing connection to host: %s\n", err)
		return
	}

	start := time.Now()
	start0 := time.Now()
	var sleep, bcount, count time.Duration

	if *rate > 0 {
		sleep = time.Duration(1e9 / *rate - *nuance)
	}

	fmt.Printf("Starting to flood %s.\n", *server)

	for {
		n, err := conn.Write(bytes)
		if err != nil {
			fmt.Println("Error sending data:", err)
		}

		if *rate > 0 {
			now := time.Now()
			diff := now.Sub(start0)
			// fmt.Printf("diff=%d\n", diff)
			if diff < sleep {
				<- time.After(sleep - diff)
			}
			start0 = time.Now()
		}

		count += 1
		bcount += time.Duration(n)

		now := time.Now()
		diff := now.Sub(start)
		if diff > 1e9 {
			fmt.Printf("bps = %d, pps = %d.\n", bcount, count)
			start = now
			bcount, count = 0, 0
		}
	}
}