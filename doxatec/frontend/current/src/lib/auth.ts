import type { User } from "$lib/interfaces"
import type { RequestEvent } from "@sveltejs/kit"

export const authenticateUser = (event: RequestEvent) => {
    // get the cookies from the request
    const { cookies } = event

    // get the user token from the cookie
    const userToken = cookies.get("auth")

    // if the user token is not valid, return null
    // this is where you would check the user token against your database
    // to see if it is valid and return the user object
    if (userToken === "demo@admin") {
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

    if (userToken === "demo@user") {
        const user: User = {
            id: "1",
            username: "user",
            password: "1234",
            avatar: "https://cdn.dribbble.com/userupload/2798814/file/original-3cfdbabadfd8f92aed97b0c0b57c6b89.png?compress=1&resize=752x",
            name: "name",
            email: "user@demo.com",
            phone: "1",
            role: "user",
            created: new Date(),
            modified: new Date(),
        }
        return user
    }

    return null
}