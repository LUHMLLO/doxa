import type { RequestEvent } from "@sveltejs/kit"

export const authenticateUser = (event: RequestEvent) => {
    // get the cookies from the request
    const { cookies } = event

    // get the user token from the cookie
    const userToken = cookies.get("auth")

    // if the user token is not valid, return null
    // this is where you would check the user token against your database
    // to see if it is valid and return the user object
    if (userToken === "demo@user") {
        const user = {
            ID: "1",
            Avatar:
                "https://cdn.dribbble.com/userupload/2798814/file/original-3cfdbabadfd8f92aed97b0c0b57c6b89.png?compress=1&resize=752x",
            Username: "ClientDev",
            Password: "demo",
        }
        return user
    }
    if (userToken === "demo@admin") {
        const user = {
            ID: "0",
            Avatar:
                "https://cdn.dribbble.com/userupload/2798814/file/original-3cfdbabadfd8f92aed97b0c0b57c6b89.png?compress=1&resize=752x",
            Username: "DoxaDev",
            Password: "demo",
        }
        return user
    }

    return null
}