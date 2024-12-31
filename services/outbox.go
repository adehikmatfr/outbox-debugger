// Package services provides service layer implementations for the Outbox Debugger application.
// This file initializes the event outbox manager and starts the cron service for processing events.
package services

import (
	"context"
	"outbox/debugger/enum"
	"time"

	"clodeo.tech/public/go-universe/pkg/db/rdbms/sqldb"
	"github.com/rs/zerolog/log"

	outbox "clodeo.tech/public/go-outbox/event_outbox"
	outboxModel "clodeo.tech/public/go-outbox/event_outbox/model"
)

// initEventOutboxManager initializes the EventOutboxManager and its database manager.
//
// Returns:
//   - EventOutboxManager: The manager responsible for handling outbox events.
//   - SqlDbManager: The SQL database manager used for outbox operations.
//
// Behavior:
//   - Connects to the database using the configuration from the `enum` package.
//   - Sets up the SQL database manager for the outbox pattern.
//   - Configures the event outbox manager with topic settings defined in `enum`.
//
// Error Handling:
//   - Logs a fatal error and exits the application if the database connection fails.
func initEventOutboxManager() (outbox.EventOutboxManager, sqldb.SqlDbManager) {
	// Step 1: Establish a connection to the SQL database.
	outboxSqldb, err := sqldb.Connect(context.Background(), sqldb.DBConfig{
		Driver:                enum.DbDriver,
		DSN:                   enum.DbDSN,
		MaxOpenConnections:    enum.DbMaxOpenConnections,
		MaxIdleConnections:    enum.DbMaxIdleConnections,
		ConnectionMaxLifetime: enum.DbConnectionMaxLifetime,
		Retry:                 enum.DbRetry,
	})

	if err != nil {
		log.Fatal().Msg(err.Error()) // Log and terminate if connection fails.
	}

	// Step 2: Create the SQL database manager for the outbox.
	sqldbOutboxManager := sqldb.New(&sqldb.Opts{
		DB: outboxSqldb,
	})

	// Step 3: Configure the topic settings for the event outbox.
	eventTopicIndexes := []outboxModel.TopicConfig{
		{
			Topic:               enum.TopicName,
			Index:               enum.TableIndex,
			DeleteExistingOnAdd: enum.DeleteExistingOnAdd,
		},
	}

	// Step 4: Return the initialized EventOutboxManager and SqlDbManager.
	return outbox.NewEventOutboxManager(sqldbOutboxManager, 5, eventTopicIndexes, false), sqldbOutboxManager
}

// StartCron starts the cron service for processing outbox events.
//
// Behavior:
//   - Initializes the EventOutboxManager using `initEventOutboxManager`.
//   - Starts the cron service with a batch size of 100 and a duration of 60 seconds.
//
// Error Handling:
//   - Logs errors encountered during the cron service initialization or runtime.
//
// Usage:
//
//	Call this function to continuously process outbox events in a background cron job.
func StartCron() {
	// Step 1: Initialize the outbox manager.
	outboxManager, _ := initEventOutboxManager()

	outboxManager.Init(nil)

	// Step 2: Start the cron service with the specified settings.
	outboxManager.StartCron(100, time.Duration(60)*time.Second)
	// Block the program from exiting.
	select {}
}
