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
	ckr.Name = "Session State"
	ckr.Exit()
}

func getSessionState(session *string) (string, error) {
	b, _ := exec.Command("query", "session", *session).Output()
	b, _ = japanese.ShiftJIS.NewDecoder().Bytes(b)

	lines := strings.Split(string(b), "\n")
	state := ""
	// Maybe there's only one line
	for _, line := range lines {
		fields := strings.Fields(line)
		state = fields[3]
		break
	}
	// If no active connections found, it may be nil, nil
	return state, nil
}

func run(args []string) *checkers.Checker {
	optSession := flag.String("session", "", "Session name")
	flag.Parse()

	state, err := getSessionState(optSession)

	if err != nil {
		return checkers.Critical(err.Error())
	}

	return checkers.NewChecker(checkers.OK, state)
}
