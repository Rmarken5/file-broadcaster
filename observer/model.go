package observer

type Observer interface {
	GetIdentifier() string
	OnUpdate([]string) error
}

type Subscriber interface {
	Subscribe(Observer)
	Unsubscribe(string)
	NotifyAll()
}
