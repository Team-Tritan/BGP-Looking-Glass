package commands

import (
	"os/exec"
	"strings"
)

func ExecuteCommand(command string) (string, error) {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func ExecuteCommandAsync(command string, resultChan chan<- string, doneChan chan<- bool) {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		resultChan <- err.Error()
	} else {
		resultChan <- string(output)
	}
	doneChan <- true
}
