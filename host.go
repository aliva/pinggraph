package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Host docs
type Host struct {
	Name string
	Host string
	User string
}

// Ping docs
func (h *Host) Ping(remote string) float64 {
	cmdStr := fmt.Sprintf(
		"ssh -o ServerAliveInterval=3 -q %s@%s ping -c 1 %s -W 1 | tail -1 | awk '{print $4}' | cut -d '/' -f 2",
		h.User,
		h.Host,
		remote,
	)
	cmd := exec.Command(
		"/bin/sh",
		"-c", cmdStr,
	)

	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return -2
	}

	output := string(outputBytes)
	output = strings.Trim(output, " \n")
	f, err := strconv.ParseFloat(output, 64)

	if err != nil {
		fmt.Println(output)
		return -1
	}

	return f
}
