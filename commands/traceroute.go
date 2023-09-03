package commands

import (
	"fmt"
	"time"
)

func ExecuteTraceroute(ip string) (string, error) {
	cmd := fmt.Sprintf("traceroute %s", ip)
	resultChan := make(chan string)
	doneChan := make(chan bool)

	go ExecuteCommandAsync(cmd, resultChan, doneChan)

	timeout := 10 * time.Second
	select {
	case response := <-resultChan:
		return response, nil
	case <-time.After(timeout):
		doneChan <- true
		return "Traceroute command timed out", nil
	}
}
