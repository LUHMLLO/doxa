import { writable } from "svelte/store";

export const sidebarState = writable<boolean>(false)