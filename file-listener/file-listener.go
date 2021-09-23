package file_listener

import (
	"fmt"
	"os"
)

type FileListener struct {
}

/*func (f *FileListener) ListenForFiles(directory string) (chan fsnotify.Event, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	defer watcher.Close()

	if err = watcher.Add(directory); err != nil {
		return nil, err
	}
	return watcher.Events, nil

}*/

func (f *FileListener) ReadDirectory(directory string) ([]string, error) {
	files := []string{}
	dirEntries, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("Error reading directory: %s - %v", directory, err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}
