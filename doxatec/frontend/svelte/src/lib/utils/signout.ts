import pb from "./pocketbase";

function signOut() {
  pb.authStore.clear();
}

export { signOut };
