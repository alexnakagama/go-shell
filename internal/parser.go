package internal

// contains code for parsing user input

import (
	"strings"
)

func ParseInput(input string) Command {
	var parts []string
	var current strings.Builder
	inQuotes := false

	for _, r := range input {
		switch r {
		case ' ':
			if inQuotes {
				current.WriteRune(r)
			} else if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		case '"':
			inQuotes = !inQuotes
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	if len(parts) == 0 {
		return Command{}
	}

	return Command{
		Name: parts[0],
		Args: parts[1:],
	}
}
