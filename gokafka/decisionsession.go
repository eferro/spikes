package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"encoding/json"
)

type Participant struct {
}

type Circle struct {
	Id           string
	Participants []Participant
}

type Command struct {
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
	CommandName string `json:"command_name"`
}

type CreateDecisionSessionCommand struct {
	Command
	Name                     string `json:"name"`
	NumParticipantsPerCircle int    `json:"num_per_circle"`
}

type DecisionSession struct {
	Name    string
	Circles []Circle
}

func NewDecisionSession() *DecisionSession {
	return &DecisionSession{}
}

func NewDecisionSessionFromHistory() *DecisionSession {
	return &DecisionSession{}
}

type DecisionSessionService struct {
}

func NewDecisionSessionService() *DecisionSessionService {
	return &DecisionSessionService{}
}

type Service interface {
	ProcessCommand(cmd interface{}) []Event
}

func (serv *DecisionSessionService) ProcessCommand(cmd interface{}) []Event {
	fmt.Println("Service", cmd)
	command := cmd.(Command)
	switch command.CommandName {
	case "CreateDecisionSession":
		var cmd CreateDecisionSessionCommand
		if err := json.Unmarshal([]byte(line), &cmd); err != nil {
			fmt.Println("Error", err)
			return []Event{}
		}
		return []Event{DecisionSessionCreated{Event{"1", cmd.ID}, fmt.Sprintf("Session_%d", 1)}}
	}
}

func main() {

	lines := make(chan string, 1)
	go func() {
		fscanner := bufio.NewScanner(os.Stdin)
		for fscanner.Scan() {
			lines <- fscanner.Text()
		}
	}()

	decissionService := NewDecisionSessionService()

	go func() {
		for line := range lines {
			if line != "" {
				var cmd Command
				if err := json.Unmarshal([]byte(line), &cmd); err != nil {
					fmt.Println("Error", err)
				} else {
					events := decissionService.ProcessCommand(cmd)
					for e := range events {
						fmt.Println("Event", e)
					}
				}
			}
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
