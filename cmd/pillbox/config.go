package main

import (
	"fmt"
	"os"

	"github.com/fteem/pbox/launchd"
)

var (
	agentConfig = &launchd.Config{
		ExecutableName:   "pillboxd",
		Program:          fmt.Sprintf("%s/bin/pillboxd", os.Getenv("GOPATH")),
		WorkingDirectory: "/Users/ie/projects/go/bin",
		Name:             "com.ieftimov.pillbox",
		DisplayName:      "Pillbox",
		LongDescription:  "Pillbox reminders agent",
		LogLocation:      "/tmp",
		RunAtLoad:        true,
		KeepAlive:        true,
		Disabled:         false,
	}
)
