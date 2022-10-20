package cli

type EventType int

const (
	_ EventType = iota
	EncJsonNotFound
	EncJsonFound
)

type EventMsg struct {
	EventType EventType
}
