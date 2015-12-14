package main

import (
	"log"
	"net"
)

type Client struct {
	conn *net.Conn
}

func main() {
	l, err := net.Listen("tcp", "localhost:3333")
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer l.Close()
	
	log.Println("Listening on localhost:3333")
	
	for {
		conn, err := l.Accept()
		
		if err != nil {
			log.Println(err)
		}
		
		go handle(conn)	
	}
}

func handle(conn net.Conn) {
	b := make([]byte, 1024)
	
	_, err := conn.Read(b)
	
	if err != nil {
		log.Println(err)
	}
	
	_, err = conn.Write([]byte("Hello from server"))
	
	if err != nil {
		log.Println(err)
	}
	
	conn.Close()
}