import { atom } from "nanostores";

export const isCartOpen = atom(false);

export function toggleCart() {
  isCartOpen.set(!isCartOpen);
}
