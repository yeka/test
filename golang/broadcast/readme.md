# Network Broadcasting using Go

Start by compiling the code:
- `GOOS=linux go build listen.go`
- `GOOS=linux go build broadcast.go` 

Then do these steps for testing:
1. Create a vitual network using docker: `docker network create -d bridge --subnet=192.168.10.0/24 mynet`
2. Then spawn 6 containers: `docker run -it --rm --network=mynet -v $(pwd):/app -w /app alpine sh`
3. Try to `./listen` in 3 containers, then `./broadcast` from another 3 containers
4. Clean up the network `docker network rm mynet`

Reference:
- https://github.com/aler9/howto-udp-broadcast-golang
