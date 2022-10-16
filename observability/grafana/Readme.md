# Grafana

Grafana is a visualisation dashboard. It can't work by itself as it doesnt store data.
There are various software that can integrate with grafana.

Traces:
  - Jaeger
  - Grafana Tempo
  - Zipkin
Metrics / Time Series:
  - Prometheus
  - Grafana Mimir
Logs:
  - Grafana Loki
  - ElasticSearch


Software used in this writing:
- Grafana v9.1.7
- Tempo v1.5.0
- Loki v2.6.1
- Mimir v2.3.1

Starting grafa server:
./bin/grafana-server


### Tempo

Starting grafana tempo:
./tempo -server.http-listen-port=16687 -storage.trace.backend=local -storage.trace.local.path=./storage -storage.trace.wal.path=wal -use-otel-tracer -search.enabled -reporting.enabled=false

-use-otel-tracer enable opentelemetry endpoint (default at :4317)
-server.http-listen-port enable http port that grafana can access

Grafana DataSource: http://127.0.0.1:16687


### Jaeger

Starting Jaeger All In One:
```
SPAN_STORAGE_TYPE=badger BADGER_EPHEMERAL=false BADGER_DIRECTORY_VALUE=./storage/badger/data BADGER_DIRECTORY_KEY=./storage/badger/key ./jaeger-all-in-one --collector.otlp.enabled --collector.grpc-server.host-port="127.0.0.1:14251" --collector.http-server.host-port=":14269" --collector.otlp.grpc.host-port=":4327" --admin.http.host-port=":14369"
```

### Loki

Loki CLI has a lot of options. The best way to start Loki is to use config file.
Example config file can be found here `loki-local-config.yaml`

Starting Loki:
```
./loki-darwin-amd64 --config.file=loki-local-config.yaml -reporting.enabled=false
```

Grafana DataSource: http://localhost:3100

Promtail is usually used to scrap logs from log file. But you can also push log directly to loki [1].

Example:
```
curl -v -H "Content-Type: application/json" -XPOST -s "http://localhost:3100/loki/api/v1/push" --data-raw   '{"streams": [{ "stream": { "foo": "bar2" }, "values": [ [ "1665836635000000000", "fizzbuzz" ] ] }]}'
```

### Mimir

Example config file can be found here `mimir-local-config.yaml`

Starting Mimir:
```
./mimir-darwin-amd64 -config.file=mimir-local-config.yaml -server.grpc-listen-port=9097
```

Grafana Datasource: http://localhost:9009/prometheus

For quick test, checkout promwrite[2] example using `http://localhost:9009/api/v1/push` as client url.



Reference:
[1] https://grafana.com/docs/loki/latest/api/#push-log-entries-to-loki
[2] https://github.com/castai/promwrite