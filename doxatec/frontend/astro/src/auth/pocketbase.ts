import PocketBase from "pocketbase";
import { atom } from "nanostores";

export const pb = new PocketBase("http://127.0.0.1:8090");
export const currentUser = atom(pb.authStore.model);

pb.authStore.onChange((auth) => {
  console.log("authStore changed", auth);
  currentUser.set(pb.authStore.model);
});

export async function signUp(name: string, pass: string, fullName: string) {
  try {
    const data = {
      username: name,
      password: pass,
      passwordConfirm: pass,
      name: fullName,
    };

    await pb.collection("users").create(data);

    await logIn(name, pass);
  } catch (err: any) {
    console.log(err.data);
  }
}

export async function logIn(name: string, pass: string) {
  await pb.collection("users").authWithPassword(name, pass);
}

export function signOut() {
  pb.authStore.clear();
}
