package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

type server struct {
}

func main() {
	s := server{}
	s.acceptClient()

}

func (s *server) acceptClient() {

	l, err := net.Listen("tcp4", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}


func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

		c.Write([]byte("Connection Established."))

}
