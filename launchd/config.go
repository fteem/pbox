package launchd

import (
	"fmt"
	"os"
	"text/template"
)

type Config struct {
	Name        string
	DisplayName string

	ExecutableName   string
	WorkingDirectory string
	LongDescription  string
	LogLocation      string
	RunAtLoad        bool
	StartInterval    int
	StartOnMount     bool
	WatchPaths       []string
	UserName         string
	GroupName        string
	InitGroups       bool
	ThrottleInterval int
	KeepAlive        bool
	Disabled         bool
	Program          string
	ProgramArguments []string
}

func (c *Config) AgentPath() string {
	return fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", os.Getenv("HOME"), c.Name)
}

func (c *Config) GenerateTemplate() *template.Template {
	return template.Must(template.New("launchdConfig").Parse(PlistTemplate()))
}
