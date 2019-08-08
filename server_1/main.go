package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"tcp/domain"
	"tcp/handler"
	"tcp/service"
)

func main() {
	dbService := service.NewDevice()
	dbhandler := handler.NewDBHandler(dbService)
	//Base := make(map[string]string)
	//srv := service.NewDevice(Base)
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

		go newConnect(conn, dbhandler)
	}
}

func newConnect(c net.Conn, h domain.DBhandler) {
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

		c.Write([]byte(h.Req(string(data))))

		//sep := strings.Split(string(data), " ")
		//fmt.Printf("read: %v \n", string(data))
		//
		//switch sep[0] {
		//case "GET":
		//	c.Write([]byte(h.GET(sep[1])))
		//	//h.GET(sep[1])
		//case "SET":
		//	h.SET(sep[1], sep[2])
		//case "KEYS":
		//	c.Write([]byte(h.KEYS()))
		//	//h.KEYS()
		//case "DEL":
		//	h.DEL(sep[1])
		//
		//}
	}

	fmt.Println("closed:\t", c.RemoteAddr())

	c.Close()
}
