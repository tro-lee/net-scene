package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9123")
	if err != nil {
		fmt.Print(err)
		return
	}

	for {
		fmt.Println("请输入:")
		var result string
		fmt.Scanln(&result)

		_, err := conn.Write([]byte(result))
		if err != nil {
			log.Fatal(err)
			continue
		}
	}
}
