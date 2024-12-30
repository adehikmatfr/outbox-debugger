// Package services provides service layer implementations for the Outbox Debugger application.
// This file defines the function for subscribing to messages from Google Cloud Pub/Sub using the Watermill library.
package services

import (
	"context"
	"outbox/debugger/enum"
	"outbox/debugger/helper"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rs/zerolog/log"
)

// SubOutboxDebugger sets up a message subscriber for the Outbox Debugger.
//
// Parameters:
//   - router: The message router responsible for handling incoming messages.
//   - logger: The Watermill logger used for logging throughout the process.
//
// Behavior:
//   - Configures a Google Cloud Pub/Sub subscriber to listen to the specified topic.
//   - Registers a no-publisher handler to process the incoming messages.
//   - Processes messages by invoking a handler function.
//
// Error Handling:
//   - Logs a fatal error and terminates the program if the subscriber creation fails.
func SubOutboxDebugger(router *message.Router, logger watermill.LoggerAdapter) {
	// Step 1: Configure the Pub/Sub subscriber
	pubSubConfig := googlecloud.SubscriberConfig{
		GenerateSubscriptionName:         func(topic string) string { return enum.SubscriberName },
		ProjectID:                        enum.ProjectId,
		DoNotCreateSubscriptionIfMissing: true,
		SubscriptionConfig: pubsub.SubscriptionConfig{
			EnableMessageOrdering: false,            // Disable message ordering
			AckDeadline:           40 * time.Second, // Set acknowledgment deadline
		},
	}

	// Step 2: Create the subscriber
	subscriber, err := googlecloud.NewSubscriber(pubSubConfig, logger)
	if err != nil {
		log.Fatal().Msgf("[OutboxDebugger] Could not create subscriber: %v", err) // Log and exit on error
	}

	// Step 3: Add a no-publisher handler to the router
	router.AddNoPublisherHandler(
		"OutboxDebugger", // Unique handler name
		enum.TopicName,   // Topic to subscribe to
		subscriber,       // Subscriber instance
		func(msg *message.Message) error {
			// Step 4: Process the message payload
			return helper.WrapProcessMessages(
				msg,
				func(ctx context.Context, payload interface{}) error {
					// Log the message payload for debugging or processing
					log.Info().Msgf("Received payload: %v", payload)

					// Acknowledge the message
					msg.Ack()

					// Return nil to indicate successful processing
					return nil
				},
				"svc.sub.OutboxDebugger", // Tracing identifier for message processing.
			)
		},
	)
}
