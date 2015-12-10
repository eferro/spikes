package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
	}

	listener := make(chan *docker.APIEvents, 10)
	defer func() {
		time.Sleep(10 * time.Millisecond)
		if err := client.RemoveEventListener(listener); err != nil {
			log.Panic(err)
		}
	}()

	err := client.AddEventListener(listener)
	if err != nil {
		log.Panic("Failed to add event listener: %s", err)
	}

	timeout := time.After(30 * time.Second)
	for {
		select {
		case msg := <-listener:
			fmt.Println("Received: %v", *msg)
		case <-timeout:
			log.Panic("%s timed out waiting on events")
		}
	}

}
