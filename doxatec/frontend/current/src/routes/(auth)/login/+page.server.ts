import { redirect } from "@sveltejs/kit"
import type { Actions } from "./$types"

export const actions: Actions = {
    default: async ({ cookies }) => {
        cookies.set("auth", "demo@user", {
            path: "/",
            httpOnly: true,
            sameSite: "strict",
            secure: process.env.NODE_ENV === "production",
            maxAge: 60 * 60 * 24 * 7, // 1 week
        })

        throw redirect(303, "/")
    },
} satisfies Actions;

// import { redirect } from "@sveltejs/kit"
// import type { Actions } from './$types';

// export const actions: Actions = {
//     login: async ({ request, fetch, cookies }) => {
//         const data = Object.fromEntries(await request.formData())

//         const res = await fetch("http://localhost:3000/auth/signin", {
//             method: "POST",
//             headers: {
//                 "Content-Type": "application/json",
//             },
//             body: JSON.stringify(data),
//         }).then((response) => response.json()).then((data) => {
//             cookies.set('jwt', data.jwt)
//             throw redirect(303, "/")

//         }).catch((error) => {
//             console.error("Error:", error);
//         });

//     },
// } satisfies Actions;