import type User from "$lib/types/user.interface";
import { writable } from "svelte/store";

export const currentUser = writable<User>({
  ID: "1234-5678-9010-1112-1314-1516-1718-1920",
  Avatar:
    "https://cdn.dribbble.com/userupload/2798814/file/original-3cfdbabadfd8f92aed97b0c0b57c6b89.png?compress=1&resize=752x",
  Username: "DevClient",
  Password: "1234567890",
  Customer: {
    ID: "2122-2223-2324-2526-2728-2930-3132-3334",
    Name: "Client de Prueba",
    Phone: "(809)-000-0000",
    Email: "client@dev.com",
  },
});
