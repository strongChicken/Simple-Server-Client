package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const port  = ":50051"

func main()  {
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", port)
	if err != nil{
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Printf("listen: %v\n", port)

	for {
		conn, err := l.Accept()
		if err != nil{
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}