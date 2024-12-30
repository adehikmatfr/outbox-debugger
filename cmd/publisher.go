// Package cmd provides command-line interface (CLI) commands for the Outbox Debugger application.
// This file defines the "publish" command for publishing messages with various configurable options.
package cmd

import (
	"fmt"
	"os"
	"outbox/debugger/services"

	"github.com/spf13/cobra"
)

var (
	// Flags for the "publish" command
	useOutbox   bool   // Indicates whether to use the outbox pattern.
	maxMsg      int    // Maximum number of messages to publish.
	orderingKey string // Ordering key for message publishing.
)

var (
	// publisherCmd defines the "publish" command for message publishing.
	publisherCmd = &cobra.Command{
		Use:   "publish",          // Command usage text.
		Short: "Publisher runner", // Brief description of the command.
		Long: `Publish messages using various options. 
Flags:
  -useOutbox        Use the outbox pattern (default: true)
  -maxMsg           Maximum number of messages to publish (default: 0)
  -orderingKey      Ordering key for messages (default: not use ordering key)`,
		Run: func(c *cobra.Command, args []string) {
			runPublisherServices() // Executes the publishing logic.
		},
	}
)

// PublisherCmd returns the "publish" command to be registered with the root command.
//
// Behavior:
//   - Defines flags for message publishing (useOutbox, maxMsg, orderingKey).
//   - Executes the runPublisherServices function when invoked.
func PublisherCmd() *cobra.Command {
	// Define flags for the publish command
	publisherCmd.Flags().BoolVar(&useOutbox, "useOutbox", true, "Use the outbox pattern")
	publisherCmd.Flags().IntVar(&maxMsg, "maxMsg", 0, "Number of messages to publish")
	publisherCmd.Flags().StringVar(&orderingKey, "orderingKey", "", "Ordering key value")
	return publisherCmd
}

// runPublisherServices executes the logic for publishing messages.
//
// Behavior:
//   - Reads configuration flags (useOutbox, maxMsg, orderingKey).
//   - Validates input flags and ensures maxMsg is greater than 0.
//   - Calls the PubOutboxDebugger function to publish messages.
//
// Error Handling:
//   - Exits with an error message if maxMsg is invalid.
//
// Output:
//   - Logs the settings and progress of the publishing process.
func runPublisherServices() {
	fmt.Printf("Running Publisher Services with settings:\n")
	fmt.Printf("  Use Outbox: %v\n", useOutbox)
	fmt.Printf("  Max Messages: %d\n", maxMsg)
	fmt.Printf("  Ordering Key: %s\n", orderingKey)

	// Validate maxMsg flag
	if maxMsg <= 0 {
		fmt.Println("Error: maxMsg must be greater than 0")
		os.Exit(1) // Exit the application with an error status.
	}

	// Publishing messages
	fmt.Println("Publishing messages...")
	services.PubOutboxDebugger(useOutbox, orderingKey, maxMsg) // Call the service to publish messages.
	fmt.Println("Done!")
}
