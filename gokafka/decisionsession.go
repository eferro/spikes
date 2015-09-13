package main

type Participant struct {
}

type Circle struct {
     Id string
     Participants []Participant
}

type DecisionSession struct {
     Name string
     Circles []Circle
}

func NewDecisionSession() *DecisionSession {
}

func NewDecisionSessionFromHistory() *DecisionSession {
}
