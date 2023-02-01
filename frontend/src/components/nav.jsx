import { Show } from "solid-js";
import { useStore } from "@nanostores/solid";
import { isCommandOpen } from "../stores/commandpalette";
import { isCartOpen } from "../stores/cart";

const Nav = (props) => {
  const $isCommandOpen = useStore(isCommandOpen);
  const $isCartOpen = useStore(isCartOpen);

  return (
    <>
      <nav class="clr-secondary flex align--center h--56p overflow-hidden noshrink ps--sticky top zindex--12 bd-blur--32 dpt-dark" style="background-color: rgba(255,255,255,0.2);">
        <container class="container block zindex--2">
          <row class="flex row gap--8 justify--center align--center">
            <a href="/" class="icon noshrink">
              <i class="iconoir-3d-select-face"></i>
            </a>

            <nav class="flex row justify--start align--center w--100 shrink gap--16 weight--medium">
              <a href="/products">Productos</a>
              <a href="/about">Nosotros</a>
              <a href="/contact">Localizanos</a>
              <a href={void 0}>Blog</a>
              <a href="/faqs">Faqs</a>
            </nav>

            <button
              class="p--0 icon"
              onClick={() => isCommandOpen.set(!$isCommandOpen())}
            >
              <i class="iconoir-search"></i>
            </button>
            <a href="/account" class="p--0 icon">
              <i class="iconoir-profile-circle"></i>
            </a>

            <i class="iconoir-minus icon rotate--90deg offx--8 "></i>

            <button
              class="p--0 icon"
              onClick={() => isCartOpen.set(!$isCartOpen())}
            >
              <i class="iconoir-simple-cart"></i>
            </button>
          </row>
        </container>
      </nav>
    </>
  );
};

export default Nav;
