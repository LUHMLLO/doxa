import type { PageServerLoad } from "./$types";
import type { Device } from "$lib/types"

export const load = (async ({ params, fetch }) => {
  const res = await fetch(`http://172.17.0.1:3000/api/master/devices/read/${params.id}`, {
    method: "GET",
    credentials: "include",
  });

  const device: Device = await res.json()

  return {
    device,
  }
}) satisfies PageServerLoad;