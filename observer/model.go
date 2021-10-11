package observer

//go:generate mockgen -destination=../mocks/mock_observer.go -package=mocks . Observer
//go:generate mockgen -destination=../mocks/mock_subscriber.go -package=mocks . Subscriber
//go:generate mockgen -destination=../mocks/mock_subject.go -package=mocks . Subject

type Observer interface {
	GetIdentifier() string
	OnUpdate([]string) error
}

type Subscriber interface {
	Subscribe(Observer)
	Unsubscribe(string)
	NotifyAll()
}

type Subject interface {
	AddFile(fileName string)
	RemoveFile(fileName string)
	Subscribe(observer Observer)
	Unsubscribe(key string)
	NotifyAll()
	SetFiles(files []string)
	GetFiles() []string
}
