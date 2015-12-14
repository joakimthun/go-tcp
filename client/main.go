package main

import (
	"log"
	"net"
	"bufio"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3333")
	
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = conn.Write([]byte("Hello?"))
	
	if err != nil {
		log.Fatal(err)
	}
	
	m, err := bufio.NewReader(conn).ReadString('\n')
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(m)
}