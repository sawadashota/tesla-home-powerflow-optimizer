import type { PageLoad } from './$types';
import { PUBLIC_API_URL } from '$env/static/public';

// since there's no dynamic data here, we can prerender
// it so that it gets served as a static asset in production
export const prerender = true;

export const load: PageLoad<ChargeSetting> = async ({ fetch, params }) => {
    const res = await fetch(`${PUBLIC_API_URL}/vehicle/charge/setting`)
    return res.json();
};
