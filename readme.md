# Outbox Debugger

Outbox Debugger is a service designed to facilitate debugging and monitoring of the outbox pattern in distributed systems. It integrates with Google Cloud Pub/Sub for message processing and provides utilities for managing outbox events with reliability and scalability.

## Features
- **Event Publishing**: Publish events using the outbox pattern to ensure reliable message delivery.
- **Message Subscription**: Subscribe to and process events from Google Cloud Pub/Sub topics.
- **Database Management**: Manage outbox records and configurations through a database abstraction layer.
- **Cron Service**: Process outbox events periodically using a cron job.

---

## Table of Contents
1. [Setup](#setup)
2. [Configuration](#configuration)
3. [Available Commands](#available-commands)
4. [Code Structure](#code-structure)
5. [Usage Examples](#usage-examples)
6. [Contributing](#contributing)
7. [License](#license)


## Setup

### Prerequisites
- Go 1.18+
- PostgreSQL
- Google Cloud Pub/Sub
- Access to Google Cloud Project

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/adehikmatfr/outbox-debugger.git
   cd outbox-debugger
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the database:
   - Create a PostgreSQL database.
   - Update the connection string in `enum/enum.go` under `DbDSN`:
     ```go
     const DbDSN = "postgres://username:password@localhost:5432/outbox_debugger?sslmode=disable"
     ```

4. Configure Google Cloud credentials:
   - Set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to point to your service account JSON file:
     ```bash
     export GOOGLE_APPLICATION_CREDENTIALS=/path/to/your/service-account.json
     ```

---

## Configuration

Key configurations are located in the `enum` package:

### Database Configuration
- `DbDriver`: Database driver (default: PostgreSQL).
- `DbDSN`: Connection string for the database.
- `DbMaxOpenConnections`, `DbMaxIdleConnections`, `DbConnectionMaxLifetime`: Database connection pooling settings.

### Pub/Sub Configuration
- `ProjectId`: Google Cloud Project ID.
- `SubscriberName`: Name of the Pub/Sub subscription.
- `TopicName`: Name of the Pub/Sub topic.

### Outbox Configuration
- `TableIndex`: Index for outbox table management.
- `DeleteExistingOnAdd`: Determines whether existing events should be deleted on add.

---

## Available Commands

The CLI commands are defined in the `cmd` package and can be executed as follows:

1. **Publish Messages**
   ```bash
   go run main.go publish --useOutbox=true --maxMsg=100 --orderingKey="example-key"
   ```
   - Publishes up to 100 messages using the outbox pattern.

2. **Listen for Messages**
   ```bash
   go run main.go listen
   ```
   - Subscribes to the Pub/Sub topic and processes incoming messages.

3. **Start Cron**
   ```bash
   go run main.go cron
   ```
   - Periodically processes outbox events using a cron job.

4. **Database Migration**
   ```bash
   go run main.go db up
   ```
   - Migrates the database to the latest version.

---

## Code Structure

### Main Packages
1. **`cmd/`**:
   - Defines CLI commands like `publish`, `listen`, `cron`, and `db`.

2. **`services/`**:
   - Implements business logic for publishing, subscribing, and cron-based event processing.

3. **`enum/`**:
   - Contains configuration constants for the database, Pub/Sub, and outbox.

4. **`helper/`**:
   - Utility functions for common operations like message processing.

---

## Usage Examples

### Publishing Messages
To publish 50 messages with an ordering key:
```bash
go run main.go publish --useOutbox=true --maxMsg=50 --orderingKey="my-key"
```

### Subscribing to Messages
To start the message subscriber:
```bash
go run main.go listen
```

### Processing Outbox Events Periodically
To start the cron job:
```bash
go run main.go cron
```

---

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`feature/your-feature-name`).
3. Commit your changes.
4. Submit a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

### Customization:
- Replace `your-repo` with the actual repository URL.
- Update `your-email@example.com` with your contact email.
- Add any additional instructions or configuration details specific to your project.
