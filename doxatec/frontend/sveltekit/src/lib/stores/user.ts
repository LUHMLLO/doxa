import type User from "$lib/types/user";
import { writable } from "svelte/store";

export const currentUser = writable<User>({
  ID: "1234-5678-9010-1112-1314-1516-1718-1920",
  Avatar:
    "https://cdn.dribbble.com/userupload/2798814/file/original-3cfdbabadfd8f92aed97b0c0b57c6b89.png?compress=1&resize=752x",
  Username: "username",
  Password: "1234567890",
  Name: "user full name",
  Phone: "(000)-000-0000",
  Email: "user@email",
});
