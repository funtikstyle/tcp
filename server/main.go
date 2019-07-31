package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

//type Database struct {
//	base map[string]string
//}

func main() {
	port := flag.Int("P", 80, "port connection")

	url := ":" + string(port)
	tcp, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	defer tcp.Close()

	fmt.Println("started ...")

	for {
		conn, err := tcp.Accept()
		if err != nil {
			log.Println(err)
		}

		go newConnect(conn)
	}
}

func newConnect(c net.Conn) {
	fmt.Println("conn:\t", c.RemoteAddr())

	r := bufio.NewReader(c)
	for {
		data, _, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return
			}
			break
		}
		sep := strings.Split(data, " ")
		//fmt.Printf("read: %v \n", string(data))
		base := make(map[string]string)

		switch sep[0] {
		case "GET":
			fmt.Println(base[sep[1]])
		case "SET":
			base[sep[1]] = sep[2]
		case "KEYS":
			for _, key := range base {
				fmt.Println(key)
			}
		case "DEL":
			delete(base, sep[1])
		}
	}

	fmt.Println("closed:\t", c.RemoteAddr())

	c.Close()
}
