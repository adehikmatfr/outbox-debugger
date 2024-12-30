// Package services provides service layer implementations for the Outbox Debugger application.
// This file defines functions for publishing messages to Google Cloud Pub/Sub using the Outbox pattern.
package services

import (
	"context"
	"database/sql"
	"fmt"
	"outbox/debugger/enum"

	outbox "clodeo.tech/public/go-outbox/event_outbox"
	"clodeo.tech/public/go-outbox/event_outbox/model"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rs/zerolog/log"
)

// PubOutboxDebugger publishes messages to Google Cloud Pub/Sub using the Outbox pattern.
//
// Parameters:
//   - orderingKey: The ordering key used for message ordering in Pub/Sub.
//   - useOutbox: The use outbox if false only publish message no save on outbox db.
//   - maxMsg: The maximum number of messages to publish.
//
// Behavior:
//   - Initializes the EventOutboxManager and Google Cloud Pub/Sub publisher.
//   - Publishes `maxMsg` number of messages using a transactional approach.
//   - Executes callback functions after successfully adding events to the Outbox.
//
// Error Handling:
//   - Logs and handles errors encountered during message publishing or transaction execution.
func PubOutboxDebugger(useOutbox bool, orderingKey string, maxMsg int) {
	// Step 1: Initialize the EventOutboxManager and SQL Database Manager.
	outboxManager, sqlDbManager := initEventOutboxManager()

	// Step 2: Configure the Google Cloud Pub/Sub publisher.
	pubSubConfig := googlecloud.PublisherConfig{
		ProjectID:                 enum.ProjectId,
		DoNotCreateTopicIfMissing: true,
		EnableMessageOrdering:     true,
		Marshaler: googlecloud.NewOrderingMarshaler(func(topic string, msg *message.Message) (string, error) {
			return msg.Metadata.Get("ordering_key"), nil
		}),
	}
	logger := watermill.NewStdLogger(false, false)

	// Step 3: Create the Pub/Sub publisher.
	publisher, err := googlecloud.NewPublisher(pubSubConfig, logger)
	if err != nil {
		log.Fatal().Msg(err.Error()) // Log and terminate if the publisher cannot be created.
	}

	// Step 4: Initialize the Outbox Manager with the publisher.
	outboxManager.Init(publisher)

	// Step 5: Publish messages to the Outbox.
	cbList := []model.AfterAddEventCallbackFunc{}
	for i := 0; i < maxMsg; i++ {
		// Wrap message publishing in a database transaction.
		if err := sqlDbManager.WrapTransaction(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
			// Construct the event message.
			msg := fmt.Sprintf("Event Message %d", i)

			// Add the message to the Outbox and get the callback function.
			cb, err := publishMessage(ctx, outboxManager, tx, useOutbox, orderingKey, msg)
			if err != nil {
				return err
			}

			// Append the callback function to the list.
			cbList = append(cbList, cb)

			return nil
		}); err != nil {
			log.Error().Msg(err.Error()) // Log errors during transaction execution.
		}
	}

	// Step 6: Execute the callback functions if any were collected.
	if len(cbList) > 0 {
		runCallbackFuncList(cbList)
	}
}

// publishMessage publishes a single message to the Outbox.
//
// Parameters:
//   - ctx: The context for the operation.
//   - outboxManager: The EventOutboxManager instance.
//   - tx: The SQL transaction in which the message is published.
//   - useOutbox: The use outbox if false only publish message no save on outbox db.
//   - orderingKey: The ordering key for the message.
//   - eventMsg: The message payload to publish.
//
// Returns:
//   - cb: A callback function to execute after adding the event.
//   - err: An error object if the operation fails.
func publishMessage(ctx context.Context, outboxManager outbox.EventOutboxManager, tx *sql.Tx, useOutbox bool, orderingKey string, eventMsg any) (cb model.AfterAddEventCallbackFunc, err error) {
	// Step 1: Construct the AddEvent object.
	msg := &model.AddEvent{
		EventTopic:   enum.TopicName,
		EventKey:     orderingKey,
		EventMessage: eventMsg,
	}

	// Step 2: Add the event to the Outbox and get the callback function.
	if useOutbox {
		cb, err = outboxManager.AddEvent(ctx, tx, msg)
		if err != nil {
			log.Error().Msgf("Error adding event to Outbox: %v", err.Error())
			return nil, err
		}
	} else {
		cb = func() {
			outboxManager.PublishEvent(ctx, msg)
		}
	}

	return cb, nil
}

// runCallbackFuncList executes a list of callback functions.
//
// Parameters:
//   - cb: A slice of callback functions to execute.
//
// Behavior:
//   - Iterates over the list and executes each callback function.
func runCallbackFuncList(cb []model.AfterAddEventCallbackFunc) {
	for _, c := range cb {
		func(f model.AfterAddEventCallbackFunc) {
			f() // Execute the callback function.
		}(c)
	}
}
