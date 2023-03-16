import { currentUser } from '$lib/stores';
import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
    await fetch("http://172.17.0.1:3000/api/auth/signout", {
        method: "POST",
        credentials: "include",
    });
    currentUser.set(null);

    throw redirect(303, "/signin")
}) satisfies PageLoad;