# Test

This is my testing ground.
It uses docker heavily for ease of setup.

## Contents

- [Fabio](/fabio) - A load balancer that integrates with Consul
- [Kubernetes](/kubernetes) - A deployment manager
- [Hashicorp Nomad](/nomad) - A workload orchestrator
- [ElasticSearch](/elasticsearch) - Search engine

## Docker Tricks

### 1. Accessing Container via IP Address

When using docker, the common usage is to use port forwarding via `-p` option.
Accessing containers via it's IP address is only possible in linux.
On Mac and Windows, docker container runs inside a some kind of virtual box.

To access container via IP on Mac, one can utilize `sshuttle` which require
to run an ssh-server container.

``` bash
docker run --rm --name sshdocker -p 2222:22 -d rastasheep/ubuntu-sshd:14.04
sshuttle -r root@127.0.0.1:2222 -N 172.17.0.0/24
```

Normally, container's IP is on `172.17.0.*`.
You can check it using `docker container inspect [container_name] | grep IPAddress`
The default root password for `rastasheep/ubuntu-sshd:14.04` is `root`.