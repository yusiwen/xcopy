package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func Send(host string, port int, debug bool) {
	message, err := io.ReadAll(os.Stdin)
	if err == nil {
		if debug {
			fmt.Println(string(message))
		}
	} else {
		log.Fatal(err)
		return
	}

	mtype := mimetype.Detect(message)
	if !mtype.Is("text/plain") {
		fmt.Println("only text messages are supported")
		return
	}

	var size int64
	size = int64(len(message))

	tcpServer, err := net.ResolveTCPAddr(TYPE, fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	err = binary.Write(conn, binary.LittleEndian, size)
	if err != nil {
		println("Write length failed:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write(message)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	fmt.Printf("send %d bytes to server\n", size)
}
