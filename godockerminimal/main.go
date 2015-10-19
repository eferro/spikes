package main

import (
       "time"

       "github.com/kr/pretty"

)


func main() {
     for {
     	 pretty.Println("Hello world")
	 time.Sleep(5 * time.Second)
     }
}