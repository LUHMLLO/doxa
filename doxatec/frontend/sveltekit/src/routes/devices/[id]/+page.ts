import { error } from "@sveltejs/kit";
import type { PageLoad } from "./$types";

export const load = (({ params }) => {
  if (params.id === "001") {
    return {
      name: "Device #01",
      temp: Math.floor(Math.random() * 32),
    };
  }

  throw error(404, "Not found");
}) satisfies PageLoad;
