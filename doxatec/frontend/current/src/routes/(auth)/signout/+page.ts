import { currentUser } from '$lib/stores';
import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
    await fetch("https://doxapi.onrender.com/api/auth/signout", {
        method: "POST",
        credentials: "include",
    });
    currentUser.set(null);

    throw redirect(303, "/signin")
}) satisfies PageLoad;