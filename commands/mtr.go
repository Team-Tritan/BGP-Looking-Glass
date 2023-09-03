package commands

import (
	"os/exec"
)

func ExecuteMTR(ip string) (string, error) {
	cmd := exec.Command("mtr", "--report", "--report-cycles", "3", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
