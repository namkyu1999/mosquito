package main

type EventType byte

const (
	_                     = iota
	EventDelete EventType = iota
	EventPut    EventType = iota
)

type Event struct {
	Sequence  uint64    // A unique record ID
	EventType EventType // The action taken
	Key       string    // The key affected by this transaction
	Value     string    // The value of a PUT the transaction
}
