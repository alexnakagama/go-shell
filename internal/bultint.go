package internal

import (
	"os"
	"strconv"
)

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
