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

	//_, _ = conn.Write([]byte("привет\n"))
	//_, _ = conn.Write([]byte("привет1\n"))
	//_, _ = conn.Write([]byte("привет2\n"))
	//_, _ = conn.Write([]byte("привет3\n"))
	//_, _ = conn.Write([]byte("привет4\n"))

	conn.Close()
}
