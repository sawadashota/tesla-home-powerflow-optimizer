import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import path from 'path';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		// https://kit.svelte.jp/docs/adapter-static
		adapter: adapter(),

		env: {
			dir: path.join('env', process.env.NODE_ENV),
		},
	}
};

export default config;
