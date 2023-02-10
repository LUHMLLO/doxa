import { atom } from "nanostores";

export const isNotificationOpen = atom(false);
export const notificationIcon = atom("notif-icon");
export const notificationTitle = atom("notification Title");
export const notificationContent = atom("notif content.");

export function toggleNotification() {
  isNotificationOpen.set(!isNotificationOpen);
}
