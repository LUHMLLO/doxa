import type { PageServerLoad } from "./$types"
import type { Device } from "$lib/types"

export const load = (async ({ fetch }) => {
    const res = await fetch("http://localhost:3000/api/devices/all")
    const devices: Array<Device> = await res.json()

    return {
        devices,
    }
}) satisfies PageServerLoad;