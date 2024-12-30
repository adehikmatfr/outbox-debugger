// Package main is the entry point for the Outbox Debugger application.
// It initializes and executes the command-line interface (CLI) for debugging and managing the outbox pattern.
package main

import (
	"outbox/debugger/cmd" // cmd package contains the CLI commands and execution logic.
)

// main serves as the entry point of the application.
// It invokes the Execute function from the cmd package to run the CLI commands.
//
// Behavior:
//   - Initializes the CLI environment.
//   - Parses command-line arguments and executes the appropriate command.
//   - Handles any errors during execution and exits gracefully.
//
// Usage:
//
//	Run the compiled binary with desired CLI flags and arguments to debug or manage the outbox.
func main() {
	cmd.Execute() // Execute the CLI commands.
}
