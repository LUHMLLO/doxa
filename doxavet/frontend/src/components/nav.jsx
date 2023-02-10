import { Show } from "solid-js";
import { useStore } from "@nanostores/solid";
import { isCommandOpen } from "../stores/commandpalette";
import { isCartOpen } from "../stores/cart";

const Nav = (props) => {
  const $isCommandOpen = useStore(isCommandOpen);
  const $isCartOpen = useStore(isCartOpen);

  return (
    <>
      <nav
        class="clr-secondary flex align--center h--56p overflow-hidden noshrink ps--sticky top zindex--12 bd-blur--32 dpt-dark"
        style="background-color: rgba(255,255,255,0.2);"
      >
        <container class="container block zindex--2">
          <row class="flex row gap--8 justify--center align--center">
            <a href="/" class="noshrink inline-flex align--center">
              <icon className="icon circle--24">
                <i class="iconoir-3d-select-face"></i>
              </icon>
            </a>

            <nav class="flex row justify--start align--center w--100 shrink gap--16 weight--medium">
              <a href="/products">Productos</a>
              <a href="/about">Nosotros</a>
              <a href={void 0}>Blog</a>
              <a href="/faqs">Faqs</a>
            </nav>

            <button
              class="inline-flex align--center p--0 bord--hidden"
              onClick={() => isCommandOpen.set(!$isCommandOpen())}
            >
              <icon className="icon circle--24">
                <i class="iconoir-search"></i>
              </icon>
            </button>

            <a href="/account" class="noshrink inline-flex align--center">
              <icon className="icon circle--24">
                <i class="iconoir-profile-circle"></i>
              </icon>
            </a>

            <icon className="icon circle--24 offx--8">
              <i class="iconoir-minus rotate--90deg  "></i>
            </icon>

            <button
              class="inline-flex align--center p--0 bord--hidden"
              onClick={() => isCartOpen.set(!$isCartOpen())}
            >
              <icon className="icon circle--24">
                <i class="iconoir-simple-cart"></i>
              </icon>
            </button>
          </row>
        </container>
      </nav>
    </>
  );
};

export default Nav;
