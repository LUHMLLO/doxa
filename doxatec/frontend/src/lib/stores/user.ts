import { atom } from "nanostores";
import pb from "../utils/pocketbase";

export const currentUser = atom(pb.authStore.model);
