package internal

// contains the core shell logic and REPL loop, which reads user input, parses it, and executes commands

import (
	"bufio"
	"fmt"
	"os"
)

func StartShell() {
	for {
		print("> ")

		var input string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		} else if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			continue
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
