package file_listener

import (
	"github.com/fsnotify/fsnotify"
	"os"
)

type IFileListener interface {
	ListenForFiles(directory string) (chan fsnotify.Event, error)
	ReadDirectory(dirEntries []os.DirEntry) ([]string, error)
}

type FileListener struct {
	Watcher *fsnotify.Watcher
}
