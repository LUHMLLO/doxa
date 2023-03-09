import type { PageServerLoad } from "./$types";
import type { Device } from "$lib/types"

export const load = (async ({ params, fetch }) => {
  const res = await fetch(`http://localhost:3000/api/v1/devices/read/${params.id}`)
  const device: Device = await res.json()

  return {
    device,
  }

}) satisfies PageServerLoad;