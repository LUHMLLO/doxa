import { Show } from "solid-js";
import { useStore } from "@nanostores/solid";
import { isCartOpen, toggleCart } from "../stores/cart";

const Cart = (props) => {
  const $isCartOpen = useStore(isCartOpen);

  return (
    <Show when={$isCartOpen()}>
      <cart class="flex column p--16 gap--0 radius-theme w--300p h--100 bg-primary dpt-dark overflow-hidden visible">
        <row class="flex row align--center">
          <h6 class="w--100 m--0">Carrito de compras</h6>

          <button class="icon p--0 circle--16 text--16  clr-muted" onClick={toggleCart}>
            <i class="iconoir-delete-circle"></i>
          </button>
        </row>

        <hr />

        <scroll class="flex column gap--16 p--0 overflow-hidden scroll--y h--100 offb--16">
          {[...Array(20).keys()].map((data) => (
            <article class="flex row align--center gap--8">
              <div class="block ratio--1-1 circle--48 bg-muted" />
              <column class="flex column w--100">
                <p class="clr-secondary weight--medium text--12 text--capitalize m--0">
                  hair removal brush
                </p>
                <small class="text--10 clr-accent weight--black">
                  DOP 369.75
                </small>
              </column>
              <button class="icon p--0 circle--16 text--16">
                <i class="iconoir-more-vert"></i>
              </button>
            </article>
          ))}

          <hr class="bg-transparent noshrink grow h--100vh m--0" />
        </scroll>

        <modal class=" bg-light dpt-dark flex column p--16 ps--absolute w--94 h--auto right bottom left mt--auto mx--auto mb--6 radius-theme overflow-hidden">
          <row class="flex row align--center">
            <p class="w--100 m--0 text--14">Promo</p>
            <a href={void 0} class="clr-tertiary w--100 text--12 text--end">
              Canjear codigo
            </a>
          </row>
          <row class="flex row align--center text--14">
            <p class="w--100 m--0 text--14">Subtotal</p>
            <small class="clr-accent weight--black noshrink">
              {369.75 * 4}
            </small>
          </row>
          <row class="flex row align--center text--14">
            <p class="w--100 m--0 text--14">ITBIS</p>
            <small class="clr-accent weight--black noshrink">21.69</small>
          </row>
          <row class="flex row align--center text--14">
            <h6 class="w--100 m--0 text--14">Total</h6>
            <small class="clr-accent weight--black noshrink">
              {1479 + 21.69}
            </small>
          </row>

          <hr class="" />

          <button class="flex justify--center clr-primary bg-dark radius-theme w--100 text--center">
            Realizar Compra
          </button>
        </modal>
      </cart>
    </Show>
  );
};

export default Cart;
