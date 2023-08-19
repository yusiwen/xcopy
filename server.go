package main

import (
	"encoding/binary"
	"fmt"
	"golang.design/x/clipboard"
	"io"
	"log"
	"net"
	"os"
)

func ServerInit() error {
	// Init returns an error if the package is not ready for use.
	return clipboard.Init()
}

func ServerStart(host string, port int) {
	listen, err := net.Listen(TYPE, fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()
	fmt.Printf("listening on %d\n", port)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleIncomingRequest(conn)
	}
}

func handleIncomingRequest(conn net.Conn) {
	defer conn.Close()

	var size int64
	err := binary.Read(conn, binary.LittleEndian, &size)
	if err != nil {
		log.Fatal("cannot read header: " + err.Error())
		return
	}

	message := make([]byte, size)
	n, err := io.ReadFull(conn, message)
	if int64(n) != size {
		log.Fatal("cannot read message")
		return
	}
	if err != nil {
		log.Fatal(err)
		return
	}

	clipboard.Write(clipboard.FmtText, message)
	log.Printf("received %d bytes to clipboard", size)
}
