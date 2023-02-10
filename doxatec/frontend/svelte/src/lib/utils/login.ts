import pb from "./pocketbase";

async function logIn(name: string, pass: string) {
  await pb.collection("users").authWithPassword(name, pass);
}

export { logIn };
