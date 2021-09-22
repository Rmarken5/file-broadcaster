package main

import (
	"fmt"
	"github.com/Rmarken5/file-broadcaster/observer"
	file_listener "github.com/Rmarken5/file-broadcaster/file-listener"
	"math/rand"
	"net"
	"time"
)

type server struct {
	FileListener file_listener.FileListener
	FileSubject observer.FileBroadcastSubject
}

func main() {
	s := server{
		FileListener: file_listener.FileListener{},
		FileSubject: observer.FileBroadcastSubject{
			Files: []string{"file1", "file2", "file3", "file4"},
			Observers: make(map[string]observer.Observer, 0),

		},
	}
	s.acceptClients()

}

func (s *server) acceptClients() {
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
	s.FileSubject.Subscribe(obs)
}

func (s *server) listenForFiles() {
	eventChannel, err := s.FileListener.ListenForFiles("/home/ryanm/programming/go/file-broadcaster/dummy")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			// watch for events
			case event := <-*eventChannel:
				fmt.Printf("EVENT! %#v\n", event)

			}
		}
	}()
}
