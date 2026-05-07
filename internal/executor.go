package internal

import "fmt"

// contains code for executing built in commands and external commands (cd, exit, ls)
// running external programs using os/exec package

func ExecuteCommand(s *Shell, cmd Command) error {
	if IsBuiltin(cmd.Name) {
		if err := ExecuteBuiltin(s, cmd.Name, cmd.Args); err != nil {
			return fmt.Errorf("builtin error: %w", err)
		}
		return nil
	}
	return cmd.Run(cmd.Name, cmd.Args...)
}
