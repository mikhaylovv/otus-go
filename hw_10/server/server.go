package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func() {
		log.Printf("Closing connection with %s", conn.RemoteAddr())
		if err := conn.Close(); err != nil {
			log.Printf("Close connection err: %v", err.Error())
		}
	}()
	_, err := conn.Write([]byte(fmt.Sprintf("Welcome to %s, friend from %s\n", conn.LocalAddr(), conn.RemoteAddr())))
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("RECEIVED: %s", text)
		if text == "quit" || text == "exit" {
			break
		}

		_, err = conn.Write([]byte(fmt.Sprintf("I have received '%s'\n", text)))
		if err != nil {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error happend on connection with %s: %v", conn.RemoteAddr(), err)
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:3302")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer func() {
		if err = l.Close(); err != nil {
			log.Fatalf("Close Listener error: %v", err.Error())
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Cannot accept: %v", err)
		}

		go handleConnection(conn)
	}
}
