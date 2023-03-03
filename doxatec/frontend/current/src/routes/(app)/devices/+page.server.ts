import type { PageServerLoad } from "./$types"
import type Device from "$lib/types/device.interface";

export const load = (async ({ fetch }) => {
    const res = await fetch("http://localhost:3000/api/devices")
    const devices: Array<Device> = await res.json()

    console.log(devices)

    return {
        devices,
    }
}) satisfies PageServerLoad;