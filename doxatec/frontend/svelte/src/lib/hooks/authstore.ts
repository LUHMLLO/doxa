import pb from "../utils/pocketbase";
import { currentUser } from "../stores/user";

pb.authStore.onChange((auth) => {
  console.log("authStore changed", auth);
  currentUser.set(pb.authStore.model);
});
