# Redis Clone in Go

A lightweight Redis-inspired in-memory key-value database written from scratch in Go.

This project implements a custom TCP server that accepts client connections, parses commands, stores data in memory, and sends responses back to clients through a simple text-based protocol.

---

# Features

- Custom TCP server built with Go's `net` package
- Multiple concurrent client connections using goroutines
- Persistent client sessions
- In-memory key-value storage
- Command parsing system
- Separate server, command, and storage layers

---

# Supported Commands

## PING

Checks if the server is running.

### Request

```text
PING
```

### Response

```text
PONG
```

---

## SET

Stores a key-value pair in memory.

### Request

```text
SET key value
```

### Example

```text
SET username alex
```

### Response

```text
OK
```

---

## GET

Retrieves a value from a key.

### Request

```text
GET key
```

### Example

```text
GET username
```

### Response

```text
alex
```

If the key does not exist:

```text
(nil)
```

---

## DEL

Deletes a key from the database.

### Request

```text
DEL key
```

### Example

```text
DEL username
```

### Response

```text
OK
```

If the key does not exist:

```text
(nil)
```

---

## QUIT

Closes the client connection.

### Request

```text
QUIT
```

### Response

```text
BYE
```

---

# Architecture

```
redis-clone-go
│
├── cmd
│   └── main.go
│
├── internal
│   │
│   ├── server
│   │   ├── server.go
│   │   
│   │
│   ├── commands
│   │   └── commands.go
│   │
│   └── store
│       └── store.go
│
├── go.mod
└── README.md
```

---

# How It Works

## TCP Server

The server listens for incoming TCP connections.

The flow is:

```
Client
   |
   |
 TCP Connection
   |
   v
Server
```

When a client connects, the server creates a new goroutine to handle that connection.

Example:

```
Client 1 ---> Goroutine 1 ---> handleConnection()

Client 2 ---> Goroutine 2 ---> handleConnection()

Client 3 ---> Goroutine 3 ---> handleConnection()
```

This allows multiple clients to communicate with the server at the same time.

---

# Command Processing

Incoming messages are parsed and executed by the command layer.

The flow:

```
Client
  |
  v
handleConnection()
  |
  v
commands.Process()
  |
  +--> PING
  +--> SET
  +--> GET
  +--> DEL
  +--> QUIT
```

The command layer is responsible for:

- Parsing user input
- Validating arguments
- Executing the correct operation
- Formatting responses

---

# Storage Layer

The database uses an in-memory hash map:

```go
map[string]string
```

Example:

```text
{
    "username": "alex",
    "language": "go"
}
```

The storage layer only manages data operations:

- Store values
- Retrieve values
- Delete values

It does not know anything about networking or commands.

---

# Concurrency

Because the server supports multiple clients, multiple goroutines may access the database at the same time.

Operations are protected depending on their behavior:

```
GET  -> RLock()
SET  -> Lock()
DEL  -> Lock()
```

Read operations can happen concurrently, while write operations require exclusive access.

---

# Requirements

- Go 1.22 or newer

Check your version:

```bash
go version
```

---

# Installation

Clone the repository:

```bash
git clone https://github.com/alexnakagama/redis-clone-go.git
```

Enter the project directory:

```bash
cd redis-clone-go
```

Build the project:

```bash
go build ./...
```

---

# Running the Server

Start the server:

```bash
go run ./cmd
```

The server will listen on:

```
localhost:6379
```

---

# Connecting to the Server

You can use `netcat` as a client:

```bash
nc localhost 6379
```

Example session:

```text
PING
PONG

SET name alex
OK

GET name
alex

DEL name
OK

QUIT
BYE
```

---

# Testing

Run all tests:

```bash
go test ./...
```

Run tests with Go race detector:

```bash
go test -race ./...
```

The race detector helps detect concurrency problems between goroutines.

---

# Project Structure Explanation

## server

Responsible for:

- Creating the TCP listener
- Accepting connections
- Managing client sessions
- Handling connection lifecycle

---

## commands

Responsible for:

- Parsing commands
- Validating arguments
- Calling storage operations
- Generating responses

---

## store

Responsible for:

- Managing in-memory data
- Providing CRUD operations
- Protecting shared data access

---

# Future Improvements

## Protocol

- [ ] Implement Redis RESP protocol
- [ ] Support Redis-compatible clients

## Commands

- [ ] EXISTS
- [ ] KEYS
- [ ] INCR
- [ ] DECR
- [ ] MGET
- [ ] MSET

## Database Features

- [ ] Data persistence
- [ ] Snapshot system
- [ ] Append-only file storage

## Server Improvements

- [ ] Configuration files
- [ ] Better logging system
- [ ] Authentication
- [ ] Benchmarking
- [ ] Improved error handling

