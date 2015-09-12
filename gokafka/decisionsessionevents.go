package main

type DecisionSessionCreated struct {
	Name         string
}

type ParticipantAdded struct {
     Name string
}

type CircleCreated struct {
     Name string
}

type ParticipantAddedToCircle struct {
     ParticipantName string
     CircleName string
}
