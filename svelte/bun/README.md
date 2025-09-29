# Bun

## Manually

Bun is an alternative to nodejs which claims to be faster.

```bash
bun init -y
bun add svelte
bun add -D vite @sveltejs/vite-plugin-svelte svelte
```

Add `./src/App.svelte`:
```svelte
<script>
    let count = 0;
    function increase() {
        count++;
    }
</script>

<h1>Count is {count}</h1>

<button onclick="{increase}">Increase</button>
```

Add `./src/main.js` or `./src/main.ts`:
```javascript
import { mount } from 'svelte'
import App from './App.svelte';

const app = mount(App, {
  target: document.body,
});

export default app;
```

Add `./index.html`:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Svelte with Bun</title>
</head>
<body>
    <script src="./src/main.js" type="module"></script>
</body>
</html>
```

Add `./vite.config.js`:
```js
import { svelte } from '@sveltejs/vite-plugin-svelte';

export default {
  plugins: [svelte()],
};
```

Add these into `./package.json`:
```js
{
  ...,
  "scripts": {
    "dev": "vite",
    "build": "vite build"
  },
  ...
}
```

To enable strict svelte 5 runes, edit `vite.config.js` and add `compilerOptions` to svelte plugin definition:
```js
export default {
    plugins: [
        svelte({
            compilerOptions: {
                runes: true
            }
        }),
    ],
};
```

## Vite

```bash
bun create vite@latest [project_name]
```

Choose svelte & typescript

To make it strict svelte 5 with runes, edit `svelte.config.js` and add `{compilerOptions: {runes: true}}` into default object:
```js
export default {
  ...,
  compilerOptions: {
    runes: true
  }
}
```
