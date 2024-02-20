package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Meta interface{}
}

type MessageMeta struct {
	Text     string
	ChatID   int
	Username string
}
