package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DefaultBuiltins() map[string]func(*Shell, []string) error {
	return map[string]func(*Shell, []string) error{
		"exit":    ExecuteExit,
		"cd":      ExecuteCd,
		"echo":    ExecuteEcho,
		"pwd":     ExecutePwd,
		"history": (*Shell).ExecuteHistory,
		"read":    ExecuteRead,
		"help":    (*Shell).ExecuteHelp,
		"source":  (*Shell).ExecuteSource,
	}
}

func ExecuteBuiltin(s *Shell, name string, args []string) error {
	return s.Builtins[name](s, args)
}

func (s *Shell) IsBuiltin(name string) bool {
	_, ok := s.Builtins[name]
	return ok
}

func ExecuteExit(s *Shell, args []string) error {
	if len(args) > 0 {
		if code, err := strconv.Atoi(args[0]); err == nil {
			os.Exit(code)
		}
	}
	os.Exit(0)
	return nil
}

func ExecuteCd(s *Shell, args []string) error {
	var dir string
	if len(args) == 0 {
		var err error
		dir, err = os.UserHomeDir()
		if err != nil {
			return err
		}
	} else {
		dir = args[0]
	}

	if err := os.Chdir(dir); err != nil {
		return err
	}
	return nil
}

func ExecuteEcho(s *Shell, args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func ExecutePwd(s *Shell, args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}

func (s *Shell) ExecuteHistory(args []string) error {
	for _, cmd := range s.History {
		fmt.Printf("%s\n", cmd)
	}
	return nil
}

func ExecuteRead(s *Shell, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("read: missing variable name")
	}
	varName := args[0]
	fmt.Print("> ")
	var input string
	fmt.Scanln(&input)
	os.Setenv(varName, input)
	return nil
}

func (s *Shell) ExecuteHelp(args []string) error {
	fmt.Println("Available built-in commands:")
	for name := range s.Builtins {
		fmt.Printf("  %s\n", name)
	}
	return nil
}

func (s *Shell) ExecuteSource(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("source: missing file name")
	}
	filename := args[0]
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		cmd := ParseInput(line)
		if cmd.Name != "" {
			if err := ExecuteCommand(s, cmd); err != nil {
				fmt.Printf("Error executing command from source: %v\n", err)
			}
		}
	}
	return nil
}
