package main

import (
	"fmt"
	"github.com/onatm/clockwerk"
	"time"
)

type DummyJob struct{}

func (d DummyJob) Run() {
	fmt.Println("Every 2 seconds")
}

func main() {
	var job DummyJob
	c := clockwerk.New()
	c.Every(2 * time.Second).Do(job)
	c.Start()
	time.Sleep(10 * time.Minute)
}
