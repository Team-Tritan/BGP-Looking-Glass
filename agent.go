package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func executeCommand(command string) (string, error) {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func executeCommandAsync(command string, resultChan chan<- string, doneChan chan<- bool) {
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

var (
	ipRegex     = regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`)
	subnetRegex = regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+\/\d+$`)
	asnRegex    = regexp.MustCompile(`^\d+$`)
)

func isValidIP(ip string) bool {
	return ipRegex.MatchString(ip)
}

func isValidSubnet(subnet string) bool {
	return subnetRegex.MatchString(subnet)
}

func isValidASN(asn string) bool {
	return asnRegex.MatchString(asn)
}

func executePing(ip string) (string, error) {
	cmd := fmt.Sprintf("ping -c 5 %s", ip)
	resultChan := make(chan string)
	doneChan := make(chan bool)

	go executeCommandAsync(cmd, resultChan, doneChan)

	timeout := 10 * time.Second
	select {
	case response := <-resultChan:
		return response, nil
	case <-time.After(timeout):
		doneChan <- true
		return "Ping command timed out", nil
	}
}

func executeBirdCommand(command string) (string, error) {
	cmd := fmt.Sprintf("sudo birdc show route %s", command)
	return executeCommand(cmd)
}

func main() {
	app := fiber.New()

	app.Get("/route", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if !isValidSubnet(ip) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid IP format")
		}
		response, err := executeBirdCommand(ip)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("Route Info for IP %s:\n%s", ip, response))
	})

	app.Get("/asn-routes", func(c *fiber.Ctx) error {
		asn := c.Query("asn")
		if !isValidASN(asn) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ASN format")
		}
		response, err := executeBirdCommand(fmt.Sprintf("where bgp_path ~ [= * %s * =] all", asn))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("BGP Routes for ASN %s:\n%s", asn, response))
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if !isValidIP(ip) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid IP format")
		}
		response, err := executePing(ip)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("Ping for IP %s:\n%s", ip, response))
	})

	log.Fatal(app.Listen(":8080"))
}
