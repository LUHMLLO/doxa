import pb from "./pocketbase";
import { logIn } from "./login";

async function signUp(name: string, pass: string, fullName: string) {
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

export { signUp };
