# SSH Tunneling

## Using `ssh` command

For example, want to connect to a redis server on internal network, but the only way to connect to it is via an ssh server.
Usually we do this using ssh tunneling.

For this example, we will setup a redis server on 192.168.10.3 and an ssh server on 192.168.10.2.
So go ahead and run `docker-compose up -d`

Unfortunately, on my mac, I cannot connect to 192.168.10.2 directly.
If we run `docker-compose ps`, we can see that while `redis` service has no port mapped to host (means we cannot access it directly from host),
the ```sshd``` service has it's internal port mapped to 2222 which we can use for tunnelling such as this:

```
ssh -L 6379:192.168.10.3:6379 -N root@127.0.0.1 -p2222
```

The root password is `root`. Now if you run `go run example1.go` and type `ping` followed by enter,
`redis` server will response with `+PONG`. Go ahead and try some redis commands.

If `ssh` is terminated, running `go run example1.go` will return an error.


## Using `go` code

Make sure `go run example1.go` return an error.
If it is, you can then try `go run example2.go`, which should connect as if using `ssh` command.
You can try `ping` and other redis command, should work the same.
Only, the tunneling now happen in `go` code.

The `example2` example works by opening a local port which a telnet library can connect to.
The connection and then forwarded to remote address via SSH tunneling.
While `example2` is running, you can now `go run example1.go` and it's also works.

The `example3` shows a more simple approach to SSH tunneling in go.
Open SSH connection, and then dial the remote address without first opening a local port.
