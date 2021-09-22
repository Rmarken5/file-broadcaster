package main

import (
	"fmt"
	"github.com/Rmarken5/file-broadcaster/observer"
	"math/rand"
	"net"
	"time"
)

type server struct {
	fileSubject observer.FileBroadcastSubject
}

func main() {
	s := server{
		fileSubject: observer.FileBroadcastSubject{
			Files: []string{"file1", "file2", "file3", "file4"},
			Observers: make(map[string]observer.Observer, 0),
		},
	}
	s.acceptClients()
}

func (s *server) acceptClients() {
	fmt.Println("Hello")
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
		go s.handleConnection(c)

	}
}

func (s *server) handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	obs := &observer.ConnectionData{
		Address: c.RemoteAddr().String(),
		Conn:    c,
	}
	fmt.Println("Addr: ", obs.GetIdentifier())
	c.Write([]byte("Connection Established.\n"))
	s.fileSubject.Subscribe(obs)
}
