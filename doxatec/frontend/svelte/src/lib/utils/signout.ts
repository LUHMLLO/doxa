import pb from "./pocketbase";
import { navigate } from "svelte-navigator";

function signOut() {
  pb.authStore.clear();
  navigate("/login", { replace: true });
}

export { signOut };
