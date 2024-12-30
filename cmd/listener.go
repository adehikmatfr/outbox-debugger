// Package cmd provides command-line interface (CLI) commands for the Outbox Debugger application.
// This file defines the "listen" command for starting listener services that process messages.
package cmd

import (
	"context"
	"outbox/debugger/services"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	// listenerCmd defines the "listen" command for starting listener services.
	listenerCmd = &cobra.Command{
		Use:   "listen",                    // Command usage text.
		Short: "listener services",         // Brief description of the command.
		Long:  "listener msg svc services", // Detailed description of the command.
		RunE:  runListenerServices,         // Function to execute when the command is run.
	}
)

// ListenerCmd returns the "listen" command to be registered with the root command.
//
// Behavior:
//   - Defines the "listen" command for processing incoming messages.
//   - Configures message routing, plugins, and middleware.
func ListenerCmd() *cobra.Command {
	return listenerCmd
}

// runListenerServices is the execution logic for the "listen" command.
//
// Parameters:
//   - cmd: The command instance triggering this function.
//   - args: Command-line arguments passed to the command.
//
// Behavior:
//   - Initializes a Watermill router with plugins and middleware.
//   - Registers handlers for processing messages using the SubOutboxDebugger function.
//   - Runs the router in a background context.
//
// Returns:
//   - nil if the router runs successfully.
//   - An error object if the router encounters an issue during initialization or execution.
func runListenerServices(cmd *cobra.Command, args []string) error {
	// Step 1: Initialize the logger for Watermill.
	logger := watermill.NewStdLogger(false, false)

	// Step 2: Create a new router for message handling.
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		log.Fatal().Msgf("could not create router: %v", err)
	}

	// Step 3: Add plugins and middleware to the router.
	router.AddPlugin(plugin.SignalsHandler) // Gracefully handles shutdown signals.
	router.AddMiddleware(
		middleware.CorrelationID, // Copies correlation ID to outgoing messages.
		middleware.Recoverer,     // Recovers from panics and passes errors to retry middleware.
	)

	// Step 4: Register message handlers.
	services.SubOutboxDebugger(router, logger)

	// Step 5: Run the router in a background context.
	ctx := context.Background()
	if err := router.Run(ctx); err != nil {
		log.Error().Msgf("Recover Event Message With Error: %v", err)
	}

	// Step 6: Return nil to indicate successful execution.
	return nil
}
