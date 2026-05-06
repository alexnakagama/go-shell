package internal

import (
	"fmt"
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

func ExecuteCd(args []string) {
	if len(args) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}
		os.Chdir(home)
		return
	}
	os.Chdir(args[0])
}
