import { authenticateUser } from "$lib/server/auth"
import { redirect, type Handle } from "@sveltejs/kit"

export const handle: Handle = (async ({ event, resolve }) => {
    //stage 1 - Request hits server : Before response is generated
    event.locals.user = authenticateUser(event)

    if (event.route.id!.startsWith("/(app)/")) {
        if (!event.locals.user) {
            throw redirect(303, "/login")
        }
    }

    //stage2 - Render route : Generate response : Await resolve(event)
    const response = await resolve(event)

    //stage3 - After response generated
    return response
}) satisfies Handle;