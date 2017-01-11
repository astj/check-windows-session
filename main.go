package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"

	"github.com/mackerelio/checkers"
	"golang.org/x/text/encoding/japanese"
)

func main() {
	ckr := run(os.Args[1:])
	ckr.Name = "Current Session"
	ckr.Exit()
}

func getCurrentSessionName(username *string) (string, error) {
	b, _ := exec.Command("query", "session", *username).Output()
	b, _ = japanese.ShiftJIS.NewDecoder().Bytes(b)

	lines := strings.Split(string(b), "\n")
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
	optUser := flag.String("user", "administrator", "User name")
	flag.Parse()

	name, err := getCurrentSessionName(optUser)

	if err != nil {
		return checkers.Critical(err.Error())
	}

	return checkers.NewChecker(checkers.OK, name)
}
