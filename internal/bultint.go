package internal

import (
	"os"
	"strconv"
)

func ExecuteExit(args []string) {
	if len(args) > 0 {
		if code, err := strconv.Atoi(args[0]); err != nil {
			os.Exit(code)
		}
	}
	os.Exit(0)
}
