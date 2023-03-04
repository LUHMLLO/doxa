import type { PageServerLoad } from "./$types";
import type { Device } from "$lib/interfaces"

export const load = (async ({ params, fetch }) => {
  const res = await fetch(`http://localhost:3000/api/devices/read/${params.id}`)
  const device: Device = await res.json()

  return {
    device,
  }

}) satisfies PageServerLoad;