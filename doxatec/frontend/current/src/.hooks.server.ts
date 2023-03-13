//import { currentUser } from "$lib/stores"
import { redirect, type Handle } from "@sveltejs/kit"

export const handle: Handle = (async ({ event, resolve }) => {

    const { cookies } = event

    const token = cookies.get("jwt")

    if (event.route.id?.startsWith("/(app)/" || event.url.pathname == "/")) {
        if (!token || token == "") {
            console.log("token it's either invalid or does not exists")
            throw redirect(303, "/signin")
        }
    }

    if (event.route.id?.startsWith("/(auth)/")) {
        if (token && token != "") {
            console.log("token it's already valid")
            throw redirect(303, "/")
        }
    }

    const response = await resolve(event);
    return response;
}) satisfies Handle