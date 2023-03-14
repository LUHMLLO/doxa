import type { LayoutServerLoad } from './$types';
import type { User } from '$lib/types';

export const load = (async ({ fetch }) => {

  const res = await fetch("https://doxapi.onrender.com/api/auth/signature", {
    method: "GET",
    credentials: "include",
  });

  const user: User = await res.json();

  return {
    user,
  };
}) satisfies LayoutServerLoad;


