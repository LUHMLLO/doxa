import { redirect } from "@sveltejs/kit"
import type { RequestHandler } from "./$types"

export const POST = (async ({ cookies }) => {
    cookies.delete("auth")
    throw redirect(303, "/")
}) satisfies RequestHandler;