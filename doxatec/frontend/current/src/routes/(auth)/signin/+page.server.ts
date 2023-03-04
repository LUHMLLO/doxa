import { redirect } from "@sveltejs/kit"
import type { Actions } from './$types';

export const actions: Actions = {
    signin: async ({ request, fetch, cookies }) => {

        try {
            const data = Object.fromEntries(await request.formData())

            const req = await fetch('http://localhost:3000/auth/signin', {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });
            if (!req.ok) {
                throw req.statusText
            }
            const res = await req.json()
            console.log(res)

            // cookies.set('jwt', data["jwt"].toString(), {
            //     path: "/",
            //     httpOnly: true,
            //     sameSite: "strict",
            //     secure: process.env.NODE_ENV === "production",
            //     maxAge: 60 * 60 * 24 * 7
            // })

            // throw redirect(303, "/")
        } catch (err) {
            console.error(err);
            return
        }
    },
} satisfies Actions;