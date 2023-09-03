package commands

import (
	"fmt"
)

func ExecuteBirdCommand(command string) (string, error) {
	cmd := fmt.Sprintf("sudo birdc show route %s", command)
	return ExecuteCommand(cmd)
}
