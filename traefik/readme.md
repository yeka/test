# Traefik Proxy

Traefik is a proxy and load balancers.

The concept is: EntryPoints -> Routers -> Services

EntryPoints is an address traefik listens to (eg: ":80", "127.0.0.1:8000").
Routers is a set of rules (checking path, host, etc) to determine where the request should be routed to.
It can also contains some middlewares to change the headers, etc, before forwarding the request to Services.
Services is the target, can be web server, any http server.

Traefik requires at least 1 Static Configuration and 1 Dynamic Configuration.

Static Configuration can be configured via cli, file ([traefik.yaml / traefik.toml](https://doc.traefik.io/traefik/getting-started/configuration-overview/#configuration-file)) or environtment variables.
It is where the EntryPoints and Dynamic Configurations are configured.
It should be rarely changes. Changing Static Configuration require Traefik to be restarted.
For the list of configuration available for Static Configuration, see the [reference](https://doc.traefik.io/traefik/reference/static-configuration/overview/) page.

Dynamic Configuration is where the Routers, middlewares and Services are defined.
It can be configured via file, docker, consul, kubernetes and other [providers](https://doc.traefik.io/traefik/providers/overview/#supported-providers) as determined in Static Configuration.


Now, we're gonna need a Traefik binary which can be downloaded from github.
In this example, we're gonna using [Traefik v2.6.3](https://github.com/traefik/traefik/releases/tag/v2.6.3).
Copy the binary to executabe path. For unix, you can put it under `/usr/local/bin`.

## Contents:
### Basic - [/basic](/basic)
A very minimal setup of Traefik as a proxy to a web server.


## Links:

    - Traefik - [https://doc.traefik.io](https://doc.traefik.io)
