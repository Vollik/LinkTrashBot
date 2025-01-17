package events

type Fetcher interface {
	Fetcher(limit int) ([]Event, error)
}

type Proccessor interface {
	Proccess(e Event) error
}
type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
