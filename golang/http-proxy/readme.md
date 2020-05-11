In this test:

## basic
- server1 - A simple HTTP server that listens on port 8001
- proxy1 - An HTTP proxy that listens on port 8002
- tcpforward - A TCP forwarder that listens on port 8003 which will forward any connection to it to server1
## goproxy
- An [proxy library](https://github.com/elazarl/goproxy) that supports HTTPS
