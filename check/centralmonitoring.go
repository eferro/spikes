package main

import (
	"time"

	"github.com/aleasoluciones/felixcheck"
)

func main() {
	checkEngine := felixcheck.NewCheckEngine(felixcheck.NewRiemannPublisher("127.0.0.1:5555"))

	checkEngine.AddCheck("google", "http", 3*time.Second, felixcheck.NewHttpChecker("https://www.google.org/", 200))
	checkEngine.AddCheck("golang", "http", 3*time.Second, felixcheck.NewHttpChecker("https://www.golang.org/", 200))

	for {
		time.Sleep(2 * time.Second)
	}
}
