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

	go Scantext(conn)

	//_, _ = conn.Write([]byte("SET aaa 1111\n"))
	//_, _ = conn.Write([]byte("SET bbb 2222\n"))
	//_, _ = conn.Write([]byte("DEL asd 1234\n"))
	//_, _ = conn.Write([]byte("KEYS\n"))
	//_, _ = conn.Write([]byte("GET aaa\n"))

	r := bufio.NewReader(conn)
	for {
		text, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return
			}
			break
		}

		//fmt.Printf("read: %v \n", string(data))
		//fmt.Fprintf(conn, text+"\n")
		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + text)
	}

	conn.Close()
}

func Scantext(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, _ = conn.Write([]byte(scanner.Text() + "\n"))
		fmt.Println(scanner.Text())
	}
}
