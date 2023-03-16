import type { PageServerLoad } from "./$types"
import type { Device } from "$lib/types"

export const load = (async ({ fetch }) => {
    const res = await fetch("http://172.17.0.1:3000/api/auth/mydevices", {
        method: "GET",
        credentials: "include",
    });

    const devices: Array<Device> = await res.json();

    return {
        devices,
    }
}) satisfies PageServerLoad;