package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	_, _ = conn.Write([]byte("SET aaa 1111\n"))
	_, _ = conn.Write([]byte("SET bbb 2222\n"))
	//_, _ = conn.Write([]byte("DEL asd 1234\n"))
	_, _ = conn.Write([]byte("KEYS\n"))
	_, _ = conn.Write([]byte("GET aaa"))

	r := bufio.NewReader(os.Stdin)
	for {
		text, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return
			}
			break
		}
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

	conn.Close()
}
