import pb from "../utils/pocketbase";
import { currentUser } from "../stores/user";

pb.authStore.onChange((auth, model) => {
  // console.log("authStore changed -- Auth: ", auth);
  // console.log("authStore changed -- Model: ", model);
  currentUser.set(model);
});
