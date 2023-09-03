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

func checkText(mimeType *mimetype.MIME) bool {
	if mimeType == nil {
		return false
	}

	fmt.Printf("MIME: %s, Parent: %s\n", mimeType, mimeType.Parent())
	if mimeType.Is("text/plain") {
		return true
	} else {
		for mtype := mimeType; mtype != nil; mtype = mtype.Parent() {
			if mtype.Is("text/plain") {
				return true
			}
		}
	}
	return false
}

func Send(host string, port int, debug bool, dryRun bool) {
	message, err := io.ReadAll(os.Stdin)
	if err == nil {
		if debug {
			fmt.Println(string(message))
		}
	} else {
		log.Fatal(err)
		return
	}

	detectedMIME := mimetype.Detect(message)
	if !checkText(detectedMIME) {
		println("Only text message is supported currently")
		os.Exit(1)
	}

	if dryRun {
		os.Exit(0)
	}

	size := int64(len(message))

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
