package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Print(err)
			break
		}

		resv := string(buf[:n])
		if resv == "EOF" {
			break
		}

		fmt.Println("resv", resv)

		conn.Write([]byte("ok"))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9123")
	if err != nil {
		fmt.Print(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}
		go process(conn)
	}
}
