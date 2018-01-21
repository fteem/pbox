package pillbox

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fteem/pbox/launchd"
)

func Install(c *launchd.Config) error {
	agentPath := c.AgentPath()

	_, err := os.Stat(agentPath)
	if err == nil {
		return fmt.Errorf("Init already exists: %s", agentPath)
	}

	file, err := os.Create(agentPath)
	if err != nil {
		return err
	}
	defer file.Close()

	template := c.GenerateTemplate()
	return template.Execute(file, &c)
}

func Uninstall(c *launchd.Config) error {
	agentPath := c.AgentPath()

	_, err := os.Stat(agentPath)
	if err != nil {
		return fmt.Errorf("File doesnt exists: %s", agentPath)
	}

	return os.Remove(agentPath)
}

func Load(c *launchd.Config) error {
	load := exec.Command("launchctl", "load", c.AgentPath())
	if err := load.Run(); err != nil {
		return err
	}
	return nil
}

func Unload(c *launchd.Config) error {
	unload := exec.Command("launchctl", "unload", c.AgentPath())
	if err := unload.Run(); err != nil {
		return err
	}
	return nil
}
