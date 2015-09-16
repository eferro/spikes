package main

import (
	"time"
)

type Timestamp time.Time

type Event struct {
	ID        string
	CommandID string
	Timestamp Timestamp
}

type DecisionSessionCreated struct {
	Event
	Name string
}

type ParticipantAdded struct {
	Event
	Name string
}

type CircleCreated struct {
	Event
	Name string
}

type ParticipantAddedToCircle struct {
	Event
	ParticipantName string
	CircleName      string
}
