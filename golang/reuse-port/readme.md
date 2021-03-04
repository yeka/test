# Reuse Port

This is an example of reusing port, so that a new server can be start before shutting down previous server. This
mechanism allows server restart without losing connections.

**How to test:**

1. Start a server `go run server.go "Hello 1"`
2. Try `go run client.go` which is basically `curl 127.0.0.1:8123` repeatedly
3. Start another server `go run server.go "Hello 2"`
4. Kill the first server


`server2.go` is basically the same with `server.go` just with more modular approach.