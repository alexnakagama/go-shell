package main

import (
	"github.com/alexnakagama/go-shell/internal"
)

func main() {
	config := internal.NewDefaultConfig()
	shell := internal.NewDefaultShell(config)
	shell.StartShell()
}
