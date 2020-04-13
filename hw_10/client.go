package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "time for trying to connect to another client")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatalln("usage: go-telnet <addr> <port> [--timeout=10s]")
	}

	dialer := &net.Dialer{
		Timeout: timeout,
	}
	conn, err := dialer.Dial("tcp", args[0] + ":" + args[1])
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	conScanner := bufio.NewScanner(conn)
	stdinScanner := bufio.NewScanner(os.Stdin)

	for {
		// read stdio
		if !stdinScanner.Scan() {
			if stdinScanner.Err() == nil {
				log.Println("stdio reads EOF")
				return
			}
			log.Fatalf("stdio scan error: %v", err)
			return
		}
		str := stdinScanner.Text()

		// send request
		if _, err := conn.Write([]byte(fmt.Sprintf("%s\n", str))); err != nil {
			log.Printf("server not responding, cannot write : %v", err)
			return
		}

		// read response
		if !conScanner.Scan() {
			if stdinScanner.Err() == nil {
				log.Printf("server sent EOF")
				return
			}
			log.Fatalf("connection scan error: %v", err)
		}
		text := conScanner.Text()
		fmt.Println(text)
	}
}
