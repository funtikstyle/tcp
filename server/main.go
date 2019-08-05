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

	Base := make(map[string]string)

	for {
		conn, err := tcp.Accept()
		if err != nil {
			log.Println(err)
		}

		go newConnect(conn, Base)
	}
}

func newConnect(c net.Conn, base map[string]string) {
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
			fmt.Println(base[sep[1]])
			c.Write([]byte(base[sep[1]]))
		case "SET":
			base[sep[1]] = sep[2]
			fmt.Println(base)
			c.Write([]byte("данные отравленны -" + sep[1] + "\n"))
		case "KEYS":
			for key, _ := range base {
				fmt.Println(key)
				c.Write([]byte(key))
			}
		case "DEL":
			delete(base, sep[1])
			fmt.Println("данные удалены")
			c.Write([]byte("данные удалены по ключу - " + sep[1] + "\n"))

		}
	}

	fmt.Println("closed:\t", c.RemoteAddr())

	c.Close()
}
