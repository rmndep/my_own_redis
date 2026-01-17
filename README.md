# my_own_redis

A small Redis-like server implemented in Go that demonstrates a simple implementation of the RESP (Redis Serialization Protocol) reader/writer and command handling.

Project structure (important files):

- `main.go` — program entry point; starts the server and accepts client connections.
- `handler.go` — command handling logic for incoming requests.
- `internal/resp/` — RESP protocol implementation:
  - `reader.go` — parsing incoming RESP values
  - `writer.go` — serializing RESP values (writes responses back to clients)
  - `value.go` — value types used by the RESP reader/writer

Features

- RESP parsing and serialization (bulk strings, simple strings, errors, arrays).
- Minimal command dispatch mechanism for handling requests from clients.
- Lightweight codebase intended for learning and experimentation.

Build

Requires Go 1.18+ (see `go.mod`). From the repository root:

```bash
go build ./...
```

Run

Run the server from the repository root:

```bash
go run main.go
```

Check `main.go` for the listener address and port used by the server. Once running, you can connect with a Redis client or plain TCP tools.

Example (using `redis-cli`)

If the server listens on `localhost:6379` you can run:

```bash
redis-cli -p 6379 PING
```

Or send a raw RESP request with `nc`:

```bash
printf "*1\r\n$4\r\nPING\r\n" | nc localhost 6379
```

Contributing

Contributions, issues, and feature requests are welcome. For larger changes, open an issue first to discuss the design.

Notes

This project is primarily educational and does not aim to be a production-grade Redis replacement. Use it to explore the RESP protocol, learn about Go network programming, and experiment with command implementations.

License

Specify a license for this repository (e.g., MIT) or add one to the project root as `LICENSE`.
