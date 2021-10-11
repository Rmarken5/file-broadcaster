package main

import (
	"flag"
	"fmt"
	file_listener "github.com/Rmarken5/file-broadcaster/file-listener"
	"github.com/Rmarken5/file-broadcaster/observer"
	"github.com/fsnotify/fsnotify"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

//go:generate mockgen -destination=../mocks/mock_net_listener.go -package=mocks net Listener

type server struct {
	FileListener file_listener.IFileListener
	FileSubject  observer.Subject
}
 var directory = flag.String("directory", "", "Directory to listen to files on.")
func main() {

	flag.Parse()

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		panic(err)
	}
	dirEntries, err := os.ReadDir(*directory)
	if err != nil {
		panic(err)
	}
	s := server{
		FileListener: &file_listener.FileListener{
			Watcher: watcher,
		},
		FileSubject: &observer.FileBroadcastSubject{
			Files:     []string{},
			Observers: make(map[string]observer.Observer, 0),
		},
	}
	done := make(chan bool)
	//directory := "/home/ryanm/programming/go/file-broadcaster/dummy"

	l, err := net.Listen("tcp4", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	go s.acceptClients(l)
	s.addFilesToSubject(dirEntries)
	go s.listenForFiles(*directory)

	for {
		select {
		case <-done:
			os.Exit(1)
		}
	}

}

func (s *server) acceptClients(listener net.Listener) {
	rand.Seed(time.Now().Unix())

	for {
		c, err := listener.Accept()
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

func (s *server) listenForFiles(directory string) error{

	fileListener := s.FileListener.ListenForFiles(directory)
	fmt.Println("listening for files.")

	done := make(chan bool)

	go func() {
		for {
			select {
			// watch for events
			case event := <- fileListener:
				fmt.Printf("EVENT! %+v\n", event)
				fileParts := strings.Split(event.Name, "/")
				fileName := fileParts[len(fileParts)-1]
				if event.Op.String() == "CREATE" {
					s.FileSubject.AddFile(fileName)
				} else if event.Op.String() == "REMOVE" {
					s.FileSubject.RemoveFile(fileName)
				}
			}
		}
	}()
	<-done
	return nil
}

func (s *server) addFilesToSubject(dirEntries []os.DirEntry) {
	files := s.FileListener.ReadDirectory(dirEntries)

	s.FileSubject.SetFiles(append(s.FileSubject.GetFiles(), files...))
}
