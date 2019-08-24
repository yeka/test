# Fabio - Basic

Fire up `docker-compose up -d`. Once it runs, you can open several tab:

1. Consul dashboard - <a href="127.0.0.1:8500" target="consului">127.0.0.1:8500</a>
2. Fabio dashboard - <a href="127.0.0.1:9998" target="fabioui">127.0.0.1:9998</a>
3. Fabio LB - <a href="127.0.0.1:9999" target="fabiolb">127.0.0.1:9999</a>
4. Example Web Server 1 - <a href="127.0.0.1:9001" target="ews1">127.0.0.1:9001</a>
5. Example Web Server 2 - <a href="127.0.0.1:9002" target="ews1">127.0.0.1:9002</a>
6. Example Web Server 3 - <a href="127.0.0.1:9003" target="ews1">127.0.0.1:9003</a>

If you already open that all, you'll notice that Fabio-LB is giving 404 error, because nothing has been configured in Fabio yet, and nothing have been registered to Consul.

Now let's start registering those web server 1-3 to Consul:
```
curl -i -X PUT -d '
{
    "ID": "web1",
    "Name": "Web",
    "Tags": ["urlprefix-/"],
    "Address": "web1",
    "Port": 80,
    "Check": {
        "Interval": "10s",
        "DeregisterCriticalServiceAfter": "90m",
        "HTTP": "http://web1/"
    }
}
' http://127.0.0.1:8500/v1/agent/service/register

curl -i -X PUT -d '
{
    "ID": "web2",
    "Name": "Web",
    "Tags": ["urlprefix-/"],
    "Address": "web2",
    "Port": 80,
    "Check": {
        "Interval": "10s",
        "DeregisterCriticalServiceAfter": "90m",
        "HTTP": "http://web2/"
    }
}
' http://127.0.0.1:8500/v1/agent/service/register

curl -i -X PUT -d '
{
    "ID": "web3",
    "Name": "Web",
    "Tags": ["urlprefix-/"],
    "Address": "web3",
    "Port": 8080,
    "Check": {
        "Interval": "10s",
        "DeregisterCriticalServiceAfter": "90m",
        "HTTP": "http://web3:8080/"
    }
}
' http://127.0.0.1:8500/v1/agent/service/register
```
If you refresh Fabio-LB now, it will show web1, web2 and web3 randomly.

You could try what happen if the service is deregistered using
```
curl -i -X PUT http://127.0.0.1:8500/v1/agent/service/deregister/web1
```

or what happen if service is down/stopped
```
docker-compose stop web1
```

You could also play around with the tags when service is registered. More info on how to manage the routing can be found on Fabio's homepage.

To register the service automatically, you could also utilize [Service Registrator](https://github.com/gliderlabs/registrator).
