package file_listener

import (
	"github.com/fsnotify/fsnotify"
	"os"
)

func (f *FileListener) ListenForFiles(directory string) chan fsnotify.Event {

	f.Watcher.Add(directory)
	return f.Watcher.Events

}

func (f *FileListener) ReadDirectory(dirEntries []os.DirEntry) []string {
	var files []string
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files
}
