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
	ID string `json:"id"`
}

type CreateDecisionSession struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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

func main() {

	lines := make(chan string, 1)
	go func() {
		fscanner := bufio.NewScanner(os.Stdin)
		for fscanner.Scan() {
			lines <- fscanner.Text()
		}
	}()

	go func() {
		for line := range lines {
			if line != "" {
				var cmd Command
				if err := json.Unmarshal([]byte(line), &cmd); err != nil {
					fmt.Println("Error", err)
				} else {
					fmt.Println("Command", cmd)
				}
			}
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
