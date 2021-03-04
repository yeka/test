# Snel

Snel is a tool/framework to compile `.svelte` component to javascript files to create web application using `deno`
and `svelte`.

From: https://github.com/crewdevio/Snel

## Basic

1. Install/download deno from https://github.com/denoland/deno/releases (it's a single binary), you can put it
   in `/usr/local/bin`
2. Install snel tools: ```deno run --allow-run https://deno.land/x/snel/install.ts```
3. Add path to snel binary by running: ```export PATH="$PATH:~/.deno/bin"```
4. Create a project: ```snel create abc```
5. New `abc` directory will be created: ```cd abc```
6. Start the project by running: ```trex run start```
7. You can now open the browser: ```http://localhost:3000/```
8. Start working on the project.
9. When done, run ```trex run build``` to build the production package. You only need the files inside `./dist/` for
   production use.

Some reading materials:
- http://getbem.com/