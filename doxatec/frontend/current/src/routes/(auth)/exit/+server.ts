import { redirect } from "@sveltejs/kit"
import type { RequestHandler } from "./$types"

export const POST: RequestHandler = (async ({ cookies }) => {
    cookies.delete("jwt")
    throw redirect(303, "/signin")
}) satisfies RequestHandler;