package observer

import "fmt"

type FileBroadcastSubject struct {
	Files     []string
	Observers map[string]Observer
}

func (f *FileBroadcastSubject) AddFile(fileName string) {
	var isExists bool
	for _, file := range f.Files {
		if fileName == file {
			isExists = true
		}
	}
	if !isExists {
		f.Files = append(f.Files, fileName)
		f.NotifyAll()
	}
}
func (f *FileBroadcastSubject) RemoveFile(fileName string) {
	newFileArr := f.Files
	for i, file := range f.Files {
		if fileName == file {
			newFileArr = append(f.Files[:i], f.Files[i+1:]...)
		}
	}
	f.Files = newFileArr
	f.NotifyAll()
}

func (f *FileBroadcastSubject) Subscribe(observer Observer) {
	fmt.Println("Adding new observer to subject: ", observer.GetIdentifier())
	f.Observers[observer.GetIdentifier()] = observer
	observer.OnUpdate(f.Files)
}

func (f *FileBroadcastSubject) Unsubscribe(key string) {
	delete(f.Observers, key)
	fmt.Printf("%s has closed their connection.\n", key)
}

func (f *FileBroadcastSubject) NotifyAll() {
	for _, obs := range f.Observers {
		if err := obs.OnUpdate(f.Files); err != nil {
			fmt.Printf("Connection closed for: %s\n", obs.GetIdentifier())
			f.Unsubscribe(obs.GetIdentifier())
		}
	}
}

func (f *FileBroadcastSubject) SetFiles(files []string)  {
	f.Files = files
}

func (f *FileBroadcastSubject) GetFiles() []string {
	return f.Files
}
