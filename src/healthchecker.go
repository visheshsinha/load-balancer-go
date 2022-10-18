package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

func startHealthCheck() {
	s := gocron.NewScheduler(time.Local)

	for _, server := range serverList {

		s.Every(2).Second().Do(server.checkHealth)
	}

	<-s.StartAsync()
}
