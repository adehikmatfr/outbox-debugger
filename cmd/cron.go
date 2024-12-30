// Package cmd provides command-line interface (CLI) commands for the Outbox Debugger application.
// This file defines the "cron" command, which starts the cron services.
package cmd

import (
	"outbox/debugger/services" // Package containing the cron service logic.

	"github.com/spf13/cobra" // Cobra library for CLI command creation.
)

var (
	// cronCmd defines the "cron" command for starting the cron services.
	cronCmd = &cobra.Command{
		Use:   "cron",                  // Command usage text.
		Short: "cron services",         // Brief description of the command.
		Long:  "cron msg svc services", // Detailed description of the command.
		RunE:  runCronServices,         // Function to execute when the command is run.
	}
)

// CronCmd returns the "cron" command to be registered with the root command.
//
// Behavior:
//   - Defines the "cron" command and associates it with the execution logic.
//   - This command starts the cron services when invoked.
func CronCmd() *cobra.Command {
	return cronCmd
}

// runCronServices is the execution logic for the "cron" command.
//
// Parameters:
//   - cmd: The command instance triggering this function.
//   - args: Command-line arguments passed to the command.
//
// Behavior:
//   - Invokes the StartCron function from the services package to initialize the cron services.
//   - Returns an error if initialization fails.
//
// Returns:
//   - nil if the services start successfully.
//   - An error object if something goes wrong during initialization.
func runCronServices(cmd *cobra.Command, args []string) error {
	// Step 1: Register and start the cron services.
	services.StartCron()

	// Step 2: Return nil to indicate successful execution.
	return nil
}
