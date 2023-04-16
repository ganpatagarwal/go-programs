package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "5050"
	connType = "tcp"
)

func main() {
	log.Println("Starting " + connType + " server on " + connHost + ":" + connPort)

	// start listener
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {

		}
	}(l)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("Error connecting:", err.Error())
			return
		}
		log.Println("Client connected.")
		log.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))
	_, err = conn.Write([]byte("received\n"))
	if err != nil {
		log.Println(err)
	}
	handleConnection(conn)
}
