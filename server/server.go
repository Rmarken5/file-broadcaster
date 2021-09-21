package main

import (
	"bufio"
	"fmt"
	"github.com/Rmarken5/file-broadcaster/observer"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	fileSubject observer.FileBroadcastSubject
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	messages := make(chan string)
	s := server{
		fileSubject: observer.FileBroadcastSubject{
			Files: []string{"file1", "file2", "file3", "file4"},
			Observers: make(map[string]observer.Observer, 0),
		},
	}
	s.acceptClients()

	go func() {
		fmt.Println("Message: ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		messages <- text
	}()

	select {
	case message := <- messages:
		fmt.Println("message: ", message)
		s.fileSubject.AddFiles(message)
	}


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
	s.fileSubject.Subscribe(obs)


}
