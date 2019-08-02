package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	_, _ = conn.Write([]byte("SET asd 1234\n"))
	_, _ = conn.Write([]byte("SET asf 1234\n"))
	_, _ = conn.Write([]byte("DEL asd 1234\n"))
	_, _ = conn.Write([]byte("KEYS\n"))
	_, _ = conn.Write([]byte("GET asf"))

	conn.Close()
}
