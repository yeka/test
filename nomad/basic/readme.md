# Nomad - Basic

From [getting started](https://www.nomadproject.io/intro/getting-started/running.html) documentation, you can run this command to test in on local:
```
sudo nomad agent -dev
```

You can check if it's running properly using:
```
nomad node status
nomad server members
```

Now try to create and start a new job (you need to have docker running):
1. `nomad job init` - create a new job file named `example.nomad`. The example is to run redis on docker.
2. `nomad job run example.nomad` - run the job file.
3. `nomad status example` - check the status of the job. Pay attention to `Allocation ID`.
4. `nomad alloc status -stats [alloc-id]` - show the status of allocated job. Here you can see the address of redis.
5. `nomad alloc logs [alloc-id]` - show the logs.
6. `http://localhost:4646/ui/` - The Web User Interface.

There's a [Hashi UI](https://github.com/jippi/hashi-ui) that provides a better UI for nomad.
```
docker run -e NOMAD_ENABLE=1 -e NOMAD_ADDR=http://host.docker.internal:4646 -p 8000:3000 jippi/hashi-ui
```

Reference:
- https://www.nomadproject.io/intro/getting-started/jobs.html