// Package cmd defines and executes the command-line interface (CLI) for the Outbox Debugger application.
// It uses Cobra to organize commands and their execution logic.
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra" // Cobra library for building CLI applications.
)

var (
	// rootCmd is the base command for the Outbox Debugger application.
	// It provides the primary entry point for all subcommands.
	rootCmd = &cobra.Command{
		Use:   "Outbox debugger Services", // Command usage description.
		Short: "Outbox debugger",          // A brief description of the root command.
		Long:  "Outbox debugger Services", // A longer description of the root command.
	}
)

// Execute initializes and runs the root command along with its subcommands.
//
// Behavior:
//   - Registers subcommands (ListenerCmd, PublisherCmd, CronCmd, DbMigrateCmd).
//   - Executes the root command based on user input.
//   - Handles any errors during execution and logs them appropriately.
//
// Usage:
//
//	Execute this function in the main package to start the CLI application.
func Execute() {
	// Step 1: Register subcommands to the root command.
	rootCmd.AddCommand(ListenerCmd())  // Register the Listener command.
	rootCmd.AddCommand(PublisherCmd()) // Register the Publisher command.
	rootCmd.AddCommand(CronCmd())      // Register the Cron command.
	rootCmd.AddCommand(DbMigrateCmd()) // Register the Database Migration command.

	// Step 2: Execute the root command.
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error()) // Log the error and terminate the program.
		os.Exit(-1)                           // Exit with a non-zero status code.
	}
}
