package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func readRoutine(ctx context.Context, conn net.Conn, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(conn)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				errChan <- errors.New("readerRoutine cannot scan msg")
				return
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
}

func writeRoutine(ctx context.Context, cancel context.CancelFunc, conn net.Conn, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				if scanner.Err() == nil {
					cancel()
					_ = conn.Close()
					return
				}
				errChan <- errors.New("writeRoutine cannot scan msg")
				return
			}

			str := scanner.Text()

			if _, err := conn.Write([]byte(fmt.Sprintf("%s\n", str))); err != nil {
				errChan <- err
				return
			}
		}
	}
}

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "time for trying to connect to another client")

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatalln("usage: go-telnet <addr> <port> [--timeout=10s]")
	}

	dialer := &net.Dialer{}
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	addr := args[0] + ":" + args[1]
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	errChan := make(chan error)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go readRoutine(ctx, conn, errChan, wg)
	go writeRoutine(ctx, cancel, conn, errChan, wg)


	select {
	case <-ctx.Done():
	case e := <-errChan:
		if e != nil {
			log.Println(e.Error())
		}
		_ = conn.Close()
		cancel()
	}

	wg.Wait()
}
