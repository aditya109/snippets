package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
		return
	}

	for {
		fmt.Println("Type a message to forward to server: ")
		var i string
		fmt.Scanln(&i)
	
		if i == "exit" {
			conn.Close()
		}
		_, err = conn.Write([]byte(i))
		if err != nil {
			log.Println(err)
			return
		}
	}
	
}
