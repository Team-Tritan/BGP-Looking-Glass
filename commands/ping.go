package commands

import (
	"fmt"
	"time"
)

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
