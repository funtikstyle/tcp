package main

import (
	"bufio"
	"flag"
	_ "flag"
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
	port := flag.String("P", "8080", "port connection")
	flag.Parse()

	tcp, err := net.Listen("tcp", ":"+*port)
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
	Base := make(map[string]string)
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
		sep := strings.Split(string(data), " ")
		//fmt.Printf("read: %v \n", string(data))

		switch sep[0] {
		case "GET":
			fmt.Println(Base[sep[1]])
		case "SET":
			Base[sep[1]] = sep[2]
			fmt.Println(Base)
		case "KEYS":
			for key, _ := range Base {
				fmt.Println(key)
			}
		case "DEL":
			delete(Base, sep[1])
		}
	}

	fmt.Println("closed:\t", c.RemoteAddr())

	c.Close()
}
