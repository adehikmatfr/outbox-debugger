// Package cmd provides the command-line interface (CLI) commands for the Outbox Debugger application.
// This file defines the "db" command for database migration operations, utilizing the golang-migrate library.
package cmd

import (
	"flag"
	"fmt"
	"os"
	"outbox/debugger/enum"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	// migrationCommand defines the "db" command for database migration.
	migrationCommand = &cobra.Command{
		Use:   "db",                 // Command usage text.
		Short: "Database Migration", // Short description of the command.
		Run: func(c *cobra.Command, args []string) {
			DatabaseMigration() // Executes the database migration logic.
		},
	}

	// flags defines command-line flags for the "db" command.
	flags = flag.NewFlagSet("db", flag.ExitOnError)

	// dir specifies the directory containing the migration files.
	dir = flags.String("dir", "migration", "directory with migration files")
)

// DbMigrateCmd returns the "db" command to be registered with the root command.
//
// Behavior:
//   - Defines the "db" command for database migration tasks.
//   - Supports subcommands like `up`, `down`, and `goto` for migration operations.
func DbMigrateCmd() *cobra.Command {
	return migrationCommand
}

var (
	// usageCommands provides detailed information on how to use the "db" command.
	usageCommands = `
Commands:
  up [N]?              Migrate all or N up migrations for the app DB.
  outbox [N]?          Migrate all or N up migrations for the outbox DB.
  goto [V]             Migrate the app DB to a specific version.
  down [N]?            Revert all or N migrations for the app DB.

For more features, visit https://github.com/golang-migrate/migrate/tree/master/cmd/migrate`
)

// DatabaseMigration executes the database migration logic.
//
// Behavior:
//   - Parses command-line arguments and executes the specified migration operation.
//   - Initializes the migration instance with the provided migration files and database connection string.
//   - Supports `up`, `down`, and `goto` operations with optional steps or version arguments.
//
// Error Handling:
//   - Logs fatal errors and terminates execution if a migration operation fails.
func DatabaseMigration() {
	flags.Usage = usage
	flags.Parse(os.Args[2:])

	args := flags.Args()

	if len(args) == 0 {
		flags.Usage()
		return
	}

	var (
		m       *migrate.Migrate
		connstr string
	)

	// Initialize the database connection string.
	connstr = enum.DbDSN

	// Create a new migration instance with the specified directory and connection string.
	m, err := migrate.New(
		fmt.Sprintf("file://db/%s", *dir),
		connstr,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize migration")
	}

	// Switch statement to handle migration commands.
	switch args[0] {
	case "up":
		if len(args) == 2 {
			step, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal().Err(err).Msg("Invalid step value for 'up' command")
			}
			err = m.Steps(step)
			if err != nil {
				log.Fatal().Err(err).Msg("Migration 'up' failed")
			}
		} else {
			err := m.Up()
			if err != nil {
				log.Fatal().Err(err).Msg("Migration 'up' failed")
			}
		}
		return
	case "down":
		if len(args) == 2 {
			step, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal().Err(err).Msg("Invalid step value for 'down' command")
			}
			err = m.Steps(step * -1)
			if err != nil {
				log.Fatal().Err(err).Msg("Migration 'down' failed")
			}
		} else {
			err := m.Down()
			if err != nil {
				log.Fatal().Err(err).Msg("Migration 'down' failed")
			}
		}
		return
	case "goto":
		if len(args) == 2 {
			step, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				log.Fatal().Err(err).Msg("Invalid version value for 'goto' command")
			}
			err = m.Migrate(uint(step))
			if err != nil {
				log.Fatal().Err(err).Msg("Migration 'goto' failed")
			}
		} else {
			usage()
		}
		return
	default:
		log.Fatal().Msg("Unknown command")
	}

	// Execute the migration "up" operation by default.
	err = m.Up()
	if err != nil {
		log.Fatal().Err(err).Msg("Default migration 'up' failed")
	}

	// Ensure the migration instance is closed properly.
	defer func() {
		sourceErr, dbErr := m.Close()
		if sourceErr != nil {
			log.Fatal().Err(sourceErr).Msg("Source close error")
		}
		if dbErr != nil {
			log.Fatal().Err(dbErr).Msg("Database close error")
		}
	}()
}

// usage displays the usage instructions for the "db" command.
func usage() {
	fmt.Println(usageCommands)
}
