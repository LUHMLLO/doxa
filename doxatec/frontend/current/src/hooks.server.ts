import { redirect, type Handle } from "@sveltejs/kit"
import { authenticateUser } from "$lib/auth"

export const handle: Handle = (async ({ event, resolve }) => {
    event.locals.user = authenticateUser(event)

    if (event.route.id?.startsWith("/(app)/") || event.url.pathname == "/") {
        if (!event.locals.user) {
            throw redirect(303, "/signin")
        }
    }

    const response = await resolve(event)

    return response
}) satisfies Handle;