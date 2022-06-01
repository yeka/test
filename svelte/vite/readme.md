# ViteJS

[ViteJS](https://vitejs.dev/) is a frontend tooling.

At the time of this writing, I use:

```
npm v8.9.0
node v18.1.0
vite v2.9.9
svelte v3.44.0
```

### Create a Project
Start a svelte project with Vite (with typescript support):

```bash
npm create vite@latest your-project -- --template svelte-ts
cd your-project
npm install
```

If you want to add scss support, simply add sass plugin:

```bash
npm add -D sass
```

Run development using:
```bash
npm run dev
```
It will start a server on localhost:3000 (default). Open the url and you'll see the default Svelte page.

### Changing Default Port

Sometimes you want to use port other than :3000.
Simply edit `vite.config.ts` and add this line `server: {port: 8080}` inside `defineConfig`. Now `vite.config.ts` should look like these:
```js
...

export default defineConfig({
  plugins: [svelte()],
  server: {
    port: 8080
  },
})

...
```

### Global SCSS/CSS

Next, add a global scss (or css) file at `src/global.scss` and put this inside those file:
```scss
main {
    h1 {
        color: blue !important;
    }
}
```
After that, edit `src/main.ts`, import the scss by adding this line:
```js
import "./global.scss"
```
It will automatically compile and update the browser. If the color if hello message becomes blue (previously red/orange), then you've successfuly add a global scss.

You can also use local scss by adding `lang="scss"` in `<style>` tag:
```html
<style lang="scss">
</style>

### Building the Project

To build the project, simply run:
```bash
npm run build
```
It will create a `./dist/` folder. Inside that folder is everything you need for production environment.