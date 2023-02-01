import { Show } from "solid-js";
import { useStore } from "@nanostores/solid";
import { isCommandOpen, toggleCommandPalette } from "../stores/commandpalette";
import clickOutside from "../utils/clickOutside";

const shortcuts = [
  {
    icon: "iconoir-compass",
    name: "Explorar",
  },
  {
    icon: "iconoir-flash",
    name: "Ofertas",
  },
  {
    icon: "iconoir-fire-flame",
    name: "Mas vendidos",
  },
  {
    icon: "iconoir-verified-badge",
    name: "Colecciones",
  },
  {
    icon: "iconoir-question-mark",
    name: "Ayuda",
  },
];

const CommandPalette = (props) => {
  const $isCommandOpen = useStore(isCommandOpen); // read the store value with the `useStore` hook

  return (
    <Show when={$isCommandOpen()}>
      <modal
        class="bg-light dpt-dark flex column mx--auto my--auto w--600p h--400p radius-theme overflow-hidden zindex--12 ps--absolute inset"
        use:clickOutside={toggleCommandPalette}
      >
        <header class="p--16">
          <row class="flex row align--center">
            <button class="icon p--0 text--16 circle--16">
              <i class="iconoir-search"></i>
            </button>

            <input
              type="text"
              placeholder="Busca productos, paginas, comandos y mas"
              class="text--12 md:text--14"
            />

            <button
              class="icon p--0 text--16 circle--16 clr-muted"
              onClick={toggleCommandPalette}
            >
              <i class="iconoir-delete-circle"></i>
            </button>
          </row>

          <hr class="noshrink bg-grayscale--0" />

          <row class="flex row justify--evenly align--center gap--16 whitespace--nowrap">
            {shortcuts.map((data) => (
              <a class="flex row align--center gap--8 noshrink m--auto">
                <icon class="icon circle--16 text--16">
                  <i class={data.icon}></i>
                </icon>
                <p class="text--12 m--0">{data.name}</p>
              </a>
            ))}
          </row>
        </header>

        <scroll class="flex column justify--start align--start overflow-hidden scroll--y h--100 mx--16 mt--0 mb--16 p--16 gap--16 radius-theme bg-primary">
          <For
            each={[
              ...Array(
                "Agendar visita veterinaria",
                "Friskies",
                "Dogshow",
                "Comida para perros",
                "Donde queda ubicado Doxavet?",
                "vacuna para gatos",
                "juguetes para mascotas",
                "Caja de arena",
                "Collar anti pulgas"
              ).values(),
            ]}
          >
            {(recentSearchResult) => (
              <a class="flex align--center gap--8 noshrink">
                <icon class="iconoir-move-right" />
                <p class="text--14 m--0">{recentSearchResult}</p>
              </a>
            )}
          </For>
        </scroll>
      </modal>
    </Show>
  );
};

export default CommandPalette;
