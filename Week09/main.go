package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalf("can't listen %s on port %d", "127.0.0.1", 9999)
	}

	reciveSize := make(chan int)

	for {
		accept, err := listen.Accept()
		if err != err {
			log.Printf("can't establish a connection by %s", err)
		}

		go func(conn net.Conn) {
			reader := bufio.NewReader(conn)
			buf := make([]byte, 1024)
			readSize, _ := reader.Read(buf)
			reciveSize <- readSize
			fmt.Println(string(buf))
		}(accept)

		go func(conn net.Conn) {
			length := <-reciveSize
			writer := bufio.NewWriter(conn)
			writer.WriteString(fmt.Sprintf("hello client, i recive size %d", length))
			writer.Flush()
		}(accept)

	}

}
