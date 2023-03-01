<script lang="ts">
  import type { PageData } from "./$types";

  export let data: PageData;
</script>

<device class="bg-light flex column p--16 theme-radius">
  <grid class="grid grid-cols--2">
    <column class="block">
      <icon class="icon"><i class="iconoir-fridge" /></icon>
      <p>{data.name}</p>
    </column>
    <column class="block text--end">
      <small>Sup: 8</small>
      <small>Mid: 4</small>
      <small>Sub: 2</small>
    </column>
  </grid>
  <row class="flex row align--center">
    <column class="flex column w--100">
      <h6>{data.temp} <sup>°</sup>C</h6>
    </column>

    <row
      class="flex row align--center gap--4 px--6 py--3 bg-primary theme-radius"
    >
      {#if data.temp >= 2 && data.temp <= 8}
        <icon class="icon size--10">
          <i class="iconoir-check-circle" />
        </icon>
        <small class="text--8">Optimo</small>
      {:else}
        <icon class="icon size--10">
          <i class="iconoir-warning-triangle" />
        </icon>
        <small class="text--8">Irregular</small>
      {/if}
    </row>
  </row>
</device>
