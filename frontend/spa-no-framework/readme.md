# Single Page Application without Framework

This is an example of how client-side routing works

Start by creating a `server.js` to serve the static html.

```bash
    docker run --name pnpm -it --init --rm -v $PWD:/app -w /app node:alpine npm init -y
    docker run --name pnpm -it --init --rm -v $PWD:/app -w /app node:alpine npm i express
    docker run --name pnpm -it --init --rm -v $PWD:/app -w /app -p 8001:8001 node:alpine node server.js
```

Reference:
- https://www.youtube.com/watch?v=6BozpmSjk-Y
- https://github.com/dcode-youtube/single-page-app-vanilla-js
