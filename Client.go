package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

const address  = ":50051"

func main()  {
	conn, err := net.Dial("tcp", address)

	if err != nil{
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("connecting to %v\n", address)

	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

func handleRead(conn net.Conn, s *sync.WaitGroup) {
	defer s.Done()

	reader, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil{
		fmt.Printf("Error of Read : %v\n", err)
		return
	}
		fmt.Printf(reader)
}

func handleWrite(conn net.Conn, s *sync.WaitGroup) {
	defer s.Done()

	ArrayData := [5]int{1, 2, 3, 4, 5}
	ByteArrayData = IntToByte(ArrayData)
	conn.Write([]byte )
	for i := 10; i > 0; i--{
		_, err := conn.Write([]byte("hello" + strconv.Itoa(i) + "\r\n"))

		if err != nil{
			fmt.Printf("error: %v\n", err)
			break
		}
	}
}

func IntToByte(n int) []byte {
	x := int(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

func ByteToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int
	binary.Read(bytesBuffer, binary.LittleEndian, &x)

	return int(x)
}
