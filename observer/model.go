package observer

//go:generate mockgen -destination=./mock_observer_test.go -package=observer . Observer
//go:generate mockgen -destination=./mock_subscriber_test.go -package=observer . Subscriber
//go:generate mockgen -destination=./mock_subject_test.go -package=observer . Subject

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
