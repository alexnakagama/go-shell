# Go Shell

A minimal, extensible shell implemented in Go, designed for learning, experimentation, and as a foundation for custom scripting environments.

## Tech Stack
- **Language:** Go (>=1.18)
- **Standard Library:** os/exec, bufio, fmt, strings, etc.
- **No external dependencies**

## Features
- REPL loop with customizable prompt
- Built-in commands: `cd`, `pwd`, `echo`, `history`, `help`, `exit`, `export`, `source`, `read`
- Command history with configurable size
- Environment variable management (`export`)
- Source external script files (`source`)
- Executes external system commands (e.g., `ls`, `cat`)
- Easily extensible with new built-in commands

## Getting Started

### Build & Run
```sh
git clone <repo-url>
cd go-shell
go run cmd/main.go
```

### Example Usage
```
> pwd
/home/user/go-shell
> echo Hello World
Hello World
> export FOO=bar
> source script.sh
> history
pwd
echo Hello World
export FOO=bar
source script.sh
history
> exit
```

## Project Structure
```
cmd/main.go         # Entry point
internal/shell.go   # Shell logic, REPL, built-ins
internal/models.go  # Command struct and parsing
internal/           # Other internal logic
```

## Adding a Built-in Command
1. Implement a function with signature: `func(*Shell, []string) error`
2. Add it to the `Builtins` map in `DefaultBuiltins()`

## Testing
Run all tests:
```sh
go test ./...
```
