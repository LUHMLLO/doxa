import type { PageServerLoad } from "./$types";
import type { Device } from "$lib/types"

export const load = (async ({ params, fetch }) => {
  const res = await fetch(`https://doxapi.onrender.com/api/master/devices/read/${params.id}`, {
    method: "GET",
    credentials: "include",
  });

  const device: Device = await res.json()

  return {
    device,
  }
}) satisfies PageServerLoad;