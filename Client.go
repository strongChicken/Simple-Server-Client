package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

const address  = ":50051"

const (
	typeString uint32 = iota // 0
	typeArray // 1
)


// length(4 bytes), type (4 bytes), body (** bytes)

func main()  {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func () {
			runClient()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("done")
}

func runClient() {
	log.Printf("connecting to %v\n", address)
	conn, err := net.Dial("tcp", address)

	if err != nil{
		log.Printf("error: %v\n", err)
		return
	}
	defer conn.Close()
	log.Println("connected", conn)

	for i := 0; i< 2; i++ {
		time.Sleep(time.Second*5)
		buffer := bytes.NewBuffer(make([]byte, 0, 1024))
		if rand.NormFloat64() > 0 {
			writeString(buffer)
		} else {
			writeString(buffer)
		}

		lenByte := make([]byte, 4)
		binary.BigEndian.PutUint32(lenByte, uint32(buffer.Len()))

		conn.Write(lenByte) // totalLength
		conn.Write(buffer.Bytes())

		log.Println("send ", buffer.Len(), "bytes")
	}
}

func writeString(buffer *bytes.Buffer) {
	// write type
	binary.Write(buffer, binary.BigEndian, typeString)

	s := "hello"
	var length uint32
	length = uint32(len(s))
	binary.Write(buffer, binary.BigEndian, length)
	binary.Write(buffer, binary.BigEndian, s)
}

func writeArray(buffer *bytes.Buffer) {
	// write type
	binary.Write(buffer, binary.BigEndian, typeArray)

	arr := []uint32{1,2,3,4,5,6,7,8}
	length := uint32(len(arr))
	binary.Write(buffer, binary.BigEndian, length)
	for _, num := range arr {
		binary.Write(buffer, binary.BigEndian, num)
	}
}

func handleRead(conn net.Conn, s *sync.WaitGroup) {
	defer s.Done()

	reader, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil{
		log.Printf("Error of Read : %v\n", err)
		return
	}
	log.Printf(reader)
}

func handleWrite(conn net.Conn, s *sync.WaitGroup) {
	defer s.Done()

	for i := 0; i < 6; i++{
		_, err := conn.Write([]byte (IntToByte(i)))

		if err != nil{
			log.Printf("error: %v\n", err)
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
