//import { currentUser } from "$lib/stores"
import { redirect, type Handle, type HandleFetch } from "@sveltejs/kit"

export const handle: Handle = (async ({ event, resolve }) => {

    const { cookies } = event

    if (event.url.pathname == "/" || event.route.id?.startsWith("/(app)/")) {
        if (!cookies.get("jwt")) {
            console.log("token not available")
        }

        const token = cookies.get("jwt")

        if (!token) {
            console.log("theres no token")
            throw redirect(303, "/signin")
        }

        if (token == "") {
            console.log("token it's not valid")
            throw redirect(303, "/signin")
        }

        const res = await fetch("http://localhost:3000/api/auth/signature", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
        });

        const data = await res.json()
        console.log(data)
    }

    const response = await resolve(event);
    return response;
}) satisfies Handle