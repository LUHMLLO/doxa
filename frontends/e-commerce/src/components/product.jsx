import {} from "solid-js";
import { useStore } from "@nanostores/solid";
import {
  isNotificationOpen,
  notificationContent,
} from "../stores/notification";

const Product = () => {
  const $isNotificationOpen = useStore(isNotificationOpen);
  const $notificationContent = useStore(notificationContent);

  function notif_AddedToCart() {
    notificationContent.set("Product was added to the cart");
    isNotificationOpen.set(!$isNotificationOpen());
    setTimeout(() => {
      isNotificationOpen.set(!$isNotificationOpen());
    }, 2500);
  }

  return (
    <>
      <article class="w--100 noshrink">
        <figure class="ps--relative overflow-hidden radius-theme dpt-grayscale--6">
          <row class="flex row justify--center align--center gap--16 w--100 ps--absolute right bottom left zindex--1 px--16 py--8">
            <button
              class="icon clr-secondary bg-light circle--24 text--14 p--0"
              onClick={() => notif_AddedToCart()}
            >
              <i class="iconoir-simple-cart" />
            </button>
            <hr class="bg-none shrink" />
            <a
              href="#"
              class="icon clr-secondary bg-light circle--24 text--14 p--0"
            >
              <i class="iconoir-move-right" />
            </a>
          </row>

          <a href="products/1" class="block ratio--1-1 w--100 p--24 bg-muted" />
        </figure>

        <p class="clr-secondary weight--medium text--capitalize offb--1">
          hair removal brush
        </p>

        <small class="text--14 clr-accent weight--black">DOP 369.75</small>
      </article>
    </>
  );
};

export default Product;
