package main

import (
	"fmt"
	"log"
	"net"
)

func testS1() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 0, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}

		log.Println("read from conn:", string(buf))
	}
}

func main() {
	fmt.Print("Hello, World!")
	testS1()
}
