package internal

// contains the core shell logic and REPL loop, which reads user input, parses it, and executes commands

import (
	"bufio"
	"fmt"
	"os"
)

type ShellConfig struct {
	MaxHistory     int
	EnableColors   bool
	WelcomeMessage string
}

type Shell struct {
	Prompt  string
	Config  ShellConfig
	History []string
}

func NewDefaultShell(config ShellConfig) *Shell {
	return &Shell{
		Prompt:  "> ",
		Config:  config,
		History: []string{},
	}
}

func NewConfig() ShellConfig {
	return ShellConfig{
		MaxHistory:     100,
		EnableColors:   true,
		WelcomeMessage: "Welcome to Go Shell!",
	}
}

func (s *Shell) StartShell() {
	fmt.Println(s.Config.WelcomeMessage)
	for {
		print(s.Prompt)
		scanner := bufio.NewScanner(os.Stdin)
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		} else if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		s.History = append(s.History, input)
		if len(s.History) > s.Config.MaxHistory {
			s.History = s.History[1:]
		}

		cmd := ParseInput(input)

		if err := ExecuteCommand(cmd); err != nil {
			fmt.Println("Error executing command:", err)
		}

		if cmd.Name == "exit" {
			break
		}
	}
}
