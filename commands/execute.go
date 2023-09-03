package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
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

func ExecutePing(ip string) (string, error) {
	cmd := fmt.Sprintf("ping -c 5 %s", ip)
	resultChan := make(chan string)
	doneChan := make(chan bool)

	go ExecuteCommandAsync(cmd, resultChan, doneChan)

	timeout := 10 * time.Second
	select {
	case response := <-resultChan:
		return response, nil
	case <-time.After(timeout):
		doneChan <- true
		return "Ping command timed out", nil
	}
}

func ExecuteBirdCommand(command string) (string, error) {
	cmd := fmt.Sprintf("sudo birdc show route %s", command)
	return ExecuteCommand(cmd)
}
