# Traefik Proxy - Basic

In this basic example, we'll start a web server on port 8100.

```shell
go run main.go
```

If you then call http://127.0.0.1:8100, it should return a "Hello world" which indicates that our web server is working.
```shell
curl http://127.0.0.1:8100
```


In this directory, we already have 2 config file. First is `traefik.yml` which is the default name for Static Configuration.
Second is `config.yml` which defined in `traefik.yml`. See the comment in each files for more information regarding the configuration.

Run the traefik, which will listen on port 8000.
```shell
traefik
```

If you call http://127.0.0.1:8000 and it returns "Hello world", the our proxy is already running successfuly.
```shell
curl http://127.0.0.1:8000
```

In this example, we enable the dashboard api and insecure mode. Therefore we can also visit [http://127.0.0.1:8080](http://127.0.0.1:8080) to open Traefik dashboard.

In the second example (traefik2.yml), we'll disable the :8080 port and configure the dashboard to listen to specific path on the port of our choice.
To run Traefik with non default Static Configuration file, use `--configFile` parameters:
```shell
traefik --configFile=traefik2.yml
```