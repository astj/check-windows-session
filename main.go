package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/mackerelio/checkers"
)

func main() {
	ckr := run(os.Args[1:])
	ckr.Name = "Windows Current Session"
	ckr.Exit()
}

func getCurrentSessionName() (string, error) {
	b, err := exec.Command("query", "session").Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(b), "\r\r\n")
	currentSessionName := ""
	for _, line := range lines {
		if strings.HasPrefix(line, ">") {
			fields := strings.Fields(line)
			currentSessionName = strings.TrimLeft(fields[0], ">")
			break
		}
	}
	// If no active connections found, it may be nil, nil
	return currentSessionName, nil
}

func run(args []string) *checkers.Checker {

	name, err := getCurrentSessionName()

	if err != nil {
		return checkers.Critical(err.Error())
	}

	return checkers.NewChecker(checkers.OK, name)
}
