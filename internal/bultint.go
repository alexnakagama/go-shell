package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var builtins = map[string]func([]string) error{
	"exit": func(args []string) error {
		ExecuteExit(args)
		return nil
	},
	"cd":   ExecuteCd,
	"echo": ExecuteEcho,
}

func IsBuiltin(name string) bool {
	_, ok := builtins[name]
	return ok
}

func ExecuteBuiltin(name string, args []string) error {
	return builtins[name](args)
}

func ExecuteExit(args []string) {
	if len(args) > 0 {
		if code, err := strconv.Atoi(args[0]); err == nil {
			os.Exit(code)
		}
	}
	os.Exit(0)
}

func ExecuteCd(args []string) error {
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

func ExecuteEcho(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
