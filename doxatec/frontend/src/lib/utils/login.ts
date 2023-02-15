import pb from "./pocketbase";
import { navigate } from "svelte-navigator";

async function logIn(name: string, pass: string) {
  await pb.collection("users").authWithPassword(name, pass);
  navigate("/dashboard", { replace: true });
}

export { logIn };
