package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"tcp/domain"
)

//var Base map[string]string

type db struct {
	srv domain.DBService
}

func NewDB(s domain.DBService) *db {
	return &db{srv: s}
}

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

func newConnect(c net.Conn, h domain.DBService) {
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
		sep := strings.Split(string(data), " ")
		//fmt.Printf("read: %v \n", string(data))

		switch sep[0] {
		case "GET":
			h.GET(sep[1])
		case "SET":
			h.SET(sep[1], sep[2])
		case "KEYS":
			h.KEYS()
		case "DEL":
			h.DEL(sep[1])

		}
	}

	fmt.Println("closed:\t", c.RemoteAddr())

	c.Close()
}
