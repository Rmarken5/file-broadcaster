package main

import (
	"fmt"
	file_listener "github.com/Rmarken5/file-broadcaster/file-listener"
	"github.com/Rmarken5/file-broadcaster/observer"
	"github.com/fsnotify/fsnotify"
	"math/rand"
	"net"
	"os"
	"time"
)

type server struct {
	FileListener file_listener.FileListener
	FileSubject  observer.FileBroadcastSubject
}

func main() {
	s := server{
		FileListener: file_listener.FileListener{},
		FileSubject: observer.FileBroadcastSubject{
			Files:     []string{},
			Observers: make(map[string]observer.Observer, 0),
		},
	}
	done := make(chan bool)
	directory := "/home/ryanm/programming/go/file-broadcaster/dummy"
	go s.acceptClients()
	s.addFilesToSubject(directory)
	go s.listenForFiles(directory)

	for {
		select {
			case <- done:
				os.Exit(1)
		}
	}
	//

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

func (s *server) listenForFiles(directory string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err.Error())
	}
	defer watcher.Close()

	if err = watcher.Add(directory); err != nil {
		panic(err)
	}
	done :=  make(chan bool)

	go func() {
		fmt.Println("listening for files.")
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %+v\n", event)
				if event.Op.String() == "CREATE" {
					s.FileSubject.AddFiles(event.Name)
				} else if event.Op.String() == "REMOVE" {
					s.FileSubject.RemoveFiles(event.Name)
				}
			}
		}
	}()
	<- done
}

func (s *server) addFilesToSubject(directory string)  {
	files, err := s.FileListener.ReadDirectory(directory)

	if err != nil {
		panic(err)
	}
	s.FileSubject.Files = append(s.FileSubject.Files, files...)
}
