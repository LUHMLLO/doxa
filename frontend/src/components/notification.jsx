import { Show } from "solid-js";
import { useStore } from "@nanostores/solid";
import {
  isNotificationOpen,
  notificationContent,
} from "../stores/notification";

const Notification = (props) => {
  const $isNotificationOpen = useStore(isNotificationOpen);
  const $notificationContent = useStore(notificationContent);

  return (
    <Show when={$isNotificationOpen()}>
      <notification class="flex row align--center gap--16 p--16 bord-complementary bord-width--2 bord--solid radius-theme w--300p h--100p bg-light dpt-dark">
        <icon class="icon">
          <i class="iconoir-bell-notification"></i>
        </icon>
        <column class="flex column w--100">
          <h6 class="m--0">notification</h6>
          <small class="m--0 text--12">{$notificationContent()}</small>
        </column>
      </notification>
    </Show>
  );
};

export default Notification;
