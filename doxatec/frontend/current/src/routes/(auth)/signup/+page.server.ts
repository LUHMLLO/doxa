import { redirect } from "@sveltejs/kit"
import type { Actions } from './$types';

export const actions: Actions = {
    signup: async ({ request, fetch }) => {
        const data = Object.fromEntries(await request.formData())

        try {
            const response = await fetch('http://localhost:3000/auth/signup', {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });
            if (!response.ok) throw response.statusText;
            throw redirect(303, "/")
        } catch (err) {
            console.error(err);
            return
        }
    },
} satisfies Actions;