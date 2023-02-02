import { atom } from "nanostores";

export const isCommandOpen = atom(false);

export function toggleCommandPalette() {
  isCommandOpen.set(!isCommandOpen);
}