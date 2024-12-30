// Package enum defines configuration constants for the Outbox Debugger application.
// These constants include database, Pub/Sub, and outbox-specific configurations.
package enum

import "clodeo.tech/public/go-universe/pkg/db/rdbms/sqldb"

// Database configuration constants.
const (
	DbDriver                = sqldb.Postgres                                                            // Database driver (PostgreSQL in this case).
	DbDSN                   = "postgres://postgres:root@localhost:5432/outbox_debbuger?sslmode=disable" // Database connection string.
	DbMaxOpenConnections    = 10                                                                        // Maximum number of open connections to the database.
	DbMaxIdleConnections    = 10                                                                        // Maximum number of idle connections in the pool.
	DbConnectionMaxLifetime = 3000                                                                      // Maximum lifetime (in seconds) of a connection.
	DbRetry                 = 3                                                                         // Number of retry attempts for database operations.
)

// Pub/Sub configuration constants.
const (
	ProjectId      = "bluebird-428713"     // Google Cloud Project ID.
	SubscriberName = "outbox.debugger-sub" // Name of the Pub/Sub subscriber.
	TopicName      = "outbox.debugger"     // Name of the Pub/Sub topic.
)

// Outbox configuration constants.
const (
	TableIndex          = 1     // Index of the table used for the outbox pattern.
	DeleteExistingOnAdd = false // Whether to delete existing entries when adding new ones.
)
