package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const port  = ":50051"

type Packet struct {
	totalLen uint32
	pkgType uint32
}

const (
	typeString uint32 = iota // 0
	typeArray // 1
)

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
		log.Println("accept")
		if err != nil{
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
		//fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	//var num int

	buf := make([]byte, 1024)
	for {
		l, err := conn.Read(buf)
		if err == io.EOF {
			// do something
			log.Println("EOF")
			return
		}
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}

		log.Println("read:", l, "bytes")

		reader := bytes.NewBuffer(buf)

		// length(4 bytes), type (4 bytes), body (** bytes)
		packet := &Packet{}

		err = binary.Read(reader, binary.BigEndian, &packet.totalLen)
		log.Println("totalLength:", packet.totalLen)

		err = binary.Read(reader, binary.BigEndian, &packet.pkgType)
		log.Println("type:", packet.pkgType)

		switch packet.pkgType {
		case typeString:
			var strLen uint32
			binary.Read(reader, binary.BigEndian, &strLen)

			var s string
			binary.Read(reader, binary.BigEndian, &s)
			log.Println("string:", s)
		case typeArray:
			var arrayLength uint32
			binary.Read(reader, binary.BigEndian, &arrayLength)

			var i uint32 = 0
			var sum uint32 = 0
			for i < arrayLength {
				var num uint32
				binary.Read(reader, binary.BigEndian, &num)
				sum += num
				i++
			}
			log.Println("sum:", sum)
		}
	}

}