import type { LayoutServerLoad } from './$types';
import type { User } from '$lib/types';

export const load = (async ({ fetch }) => {

  const res = await fetch("http://172.17.0.1:3000/api/auth/signature", {
    method: "GET",
    credentials: "include",
  });

  const user: User = await res.json();

  return {
    user,
  };
}) satisfies LayoutServerLoad;


