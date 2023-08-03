# ScyllaDB

[ScyllaDB](https://www.scylladb.com/) is a NoSQL database.


Start playing with scylla bu running:

```sh
# Start a Scylla instance
docker compose up -d

# Check status
docker compose exec scylla nodetool status

# Run Cqlsh client (https://www.tutorialspoint.com/cassandra/cassandra_cqlsh.htm)
docker compose exec -it scylla nodetool status
```
