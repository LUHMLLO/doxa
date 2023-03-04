import { writable } from "svelte/store";
import type { User } from "$lib/interfaces"

export const sidebarState = writable<boolean>(false)
export const searchModalState = writable<boolean>(false)
export const notificationState = writable<boolean>(false)

export const currentUser = writable<User>()