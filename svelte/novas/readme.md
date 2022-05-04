# Novas

Novas is a build tool that lets developers easily set up Svelte application in Deno.

From: https://github.com/NOVASland/NOVAS

## Basic

1. Install/download deno from https://github.com/denoland/deno/releases (it's a single binary), you can put it
   in `/usr/local/bin`
2. Install Novas: ```deno install --allow-net --allow-read --allow-write --unstable https://deno.land/x/novas/cli.ts```
3. Add path to snel binary by running: ```export PATH="$PATH:~/.deno/bin"```
4. Create a project: ```novas create abc```
5. New `abc` directory will be created: ```cd abc```
6. Build the project: ```novas build```; and then run development server: ```novas dev```
7. You can now open the browser: ```http://localhost:3000/```
8. Start working on the project.

At this time of writing, there's no mechanism to build for production.

Reference:

- https://medium.com/codex/novas-accelerating-svelte-and-deno-application-generation-3371c395461a