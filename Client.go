package main

import (
	"bufio"
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

	reader := bufio.NewReader(conn)
	for i := 1; i < 10; i++{
		line , err := reader.ReadString(byte('\n'))
		if err != nil{
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Printf(line)
	}
}

func handleWrite(conn net.Conn, s *sync.WaitGroup) {
	defer s.Done()

	for i := 10; i > 0; i--{
		_, err := conn.Write([]byte("hello" + strconv.Itoa(i) + "\r\n"))

		if err != nil{
			fmt.Printf("error: %v\n", err)
			break
		}
	}
}
