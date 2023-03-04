import type { RequestEvent } from "@sveltejs/kit"
import type { User } from "$lib/interfaces"

export const authenticateUser = (event: RequestEvent) => {
    const { cookies } = event

    if (cookies.get("jwt")) {
        console.log(cookies.get("jwt"))
        const user: User = {
            id: "1",
            username: "admin",
            password: "1234",
            avatar: "https://cdn.dribbble.com/userupload/2798814/file/original-3cfdbabadfd8f92aed97b0c0b57c6b89.png?compress=1&resize=752x",
            name: "name",
            email: "admin@demo.com",
            phone: "1",
            role: "admin",
            created: new Date(),
            modified: new Date(),
        }
        return user
    }

    return null
}