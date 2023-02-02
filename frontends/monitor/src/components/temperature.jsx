import { Show } from "solid-js";

const Temperature = (props) => {
  return (
    <>
      <widget class="block bg-complementary dpt-dark w--100 p--16 radius-theme overflow-hidden">
        <row class="flex row align--center">
          <h6 class="clr-tertiary m--0 w--100 text--14">Temperature</h6>

          <Show when={props.temp >= 18}>
            <span class="inline-flex justify--center align--center gap--4 clr-tertiary text--12 m--0 px--12 py--6 bg-primary radius-theme noshrink">
              <icon class="icon circle--12 text--12 p--8">
                <i class="iconoir-warning-triangle" />
              </icon>
              <small class="text--12 weight--medium">Unhealthy</small>
            </span>
          </Show>

          <Show when={props.temp <= 18}>
            <span class="inline-flex justify--center align--center gap--4 clr-tertiary text--12 m--0 px--12 py--6 bg-primary radius-theme noshrink">
              <icon class="icon circle--12 text--12 p--8">
                <i class="iconoir-lens" />
              </icon>
              <small class="text--12 weight--medium">Normal</small>
            </span>
          </Show>
        </row>
        <h3 class="offt--8 mb--0">{props.temp}</h3>
        <p class="m--0 text--12 weight--medium">{props.name}</p>
      </widget>
    </>
  );
};

export default Temperature;
