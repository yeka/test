# ViteJS

[ViteJS](https://vitejs.dev/) is a frontend tooling.

At the time of this writing, I use:

```
npm v8.9.0
node v18.1.0
vite v2.9.9
svelte v3.44.0
```

## Create a Project

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

## Changing Default Port

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

## Path Aliasing

When importing a file from sub folders, using '../..' can be dauting.
Instead, use '@' to refer to './src/' folder by using path alias.

Install some dependencies:
```bash
npm add -D path @types/node
```

Edit `vite.config.ts`:
```js
...
import path from 'path'
...

export default defineConfig((env) => ({
  ...
  resolve:{
    alias:{
      '@' : path.resolve(__dirname, './src/')
    },
  },
  ...
}))
```

Also edit `tsconfig.json`:
```
{
  ...
  "compilerOptions": {
    ...
    "paths": {
      "@/*": ["./src/*"],
    }
    ...
  },
  ...
}
```

## Global SCSS/CSS

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
```

## Adding Bootstrap

If you want to work with bootstrap, simply run:
```bash
npm add -D bootstrap
```
Add this on your global scss:
```css
@import "bootstrap";
```

To test it, add this line into `App.svelte`:
```html
<div class="text-success bg-warning">Bootstrap</div>
```

Check the browser, if you see a green text within yellow box, it means the `bootstrap` is active.

## Building the Project

To build the project, simply run:
```bash
npm run build
```
It will create a `./dist/` folder. Inside that folder is everything you need for production environment.

## Reducing Bundle Size

If you add bootstrap, you'll notice the final css size is become quite large. If you only use some of the bootstrap feature, you can purge unused css from the final bundle. Start by adding PostCSS & PurgeCSS:

```bash
npm add -D postcss "@fullhuman/postcss-purgecss"
```

Then add a `postcss.config.cjs` file with this content:
```js
const purgecss = require("@fullhuman/postcss-purgecss")

module.exports = ({ env }) => {
    const postcss = {
        plugins: []
    }

    if (env === "production") {
        postcss.plugins.push(purgecss({
            content: ['index.html', '**/*.js', '**/*.ts', '**/*.html', '**/*.svelte']
        }))
    }

    return postcss
}
```

Now if you rebuild the package, you'll notifce the final css has been reduced significantly.

## Adding Tailwind

So you wanna add tailwindcss? Let's start by adding the npm package:
```bash
npm add -D tailwindcss
```

Create a file `tailwind.config.cjs` with these content:
```js
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

Update `postcss.config.cjs`, add `require("tailwindcss"),` inside `plugins: []`. 
It should looks like this:
```js
const purgecss = require("@fullhuman/postcss-purgecss")

module.exports = ({ env }) => {
    const postcss = {
        plugins: [
            require("tailwindcss"),
        ]
    }

    if (env === "production") {
        postcss.plugins.push(purgecss({
            content: ['index.html', '**/*.js', '**/*.ts', '**/*.html', '**/*.svelte']
        }))
    }

    return postcss
}
```

Update `src/global.scss` and includes 3 new lines:
```css
@import "bootstrap"; // from previous example
@tailwind base;
@tailwind components;
@tailwind utilities;
```

Start using it in your `src/App.svelte`:
```html
<div class="bg-clip-text text-center bg-gradient-to-r from-pink-500 to-violet-500">Hello world</div>
```

Please note that adding tailwindcss will change your global css. The svelte logo will appear on the left (normally center).

Tailwind usually works alongside `autoprefixer`, so you might also wanna add that using `npm add -D autoprefixer`, and then add `require("autoprefixer")` in `postcss.config.cjs`.

## Icons

There are cool icons you can freely use in your projects.
The biggest collection of them all is Iconify ([https://iconify.design](https://iconify.design)).
It has a plugin for svelte (`@iconify/svelte`) but the bundle is quite large.

Alternatively, you can use mine. Check it out at [https://github.com/yeka/iconify](https://github.com/yeka/iconify).

Quick setup:
```bash
npm i -D '@iconify/json' 'git+https://github.com/yeka/iconify.git'
```

`vite.config.js`
```js
import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import yekaIconify from '@yeka/iconify/vite.js'

export default defineConfig((env) => ({
    plugins: [
        yekaIconify(env),
        svelte(),
    ]
}))
```

`App.svelte`
```html
<script>
    import Icon from '@yeka/iconify/Icon.svelte'
</script>
<Icon icon="mdi:home" />
```

Search any icon here [https://icon-sets.iconify.design/](https://icon-sets.iconify.design/) or here [https://icones.js.org/](https://icones.js.org/).
Icones also has a Visual Studio Code plugin.
