<script lang="ts">
  export let id: string;
  export let owner: string;
  export let name: string;
  export let temp_sup: number;
  export let temp_mid: number;
  export let temp_sub: number;
</script>

<a
  href="/devices/{id}"
  class="bg-light clr-secondary flex w--100 column p--16 theme-radius"
>
  <grid class="grid grid-cols--2">
    <column class="block">
      <icon class="icon"><i class="iconoir-fridge" /></icon>
      <p>{name}</p>
    </column>
    <column class="block text--end">
      <small>Sup: {temp_sup}</small>
      <small>Mid: {temp_mid}</small>
      <small>Sub: {temp_sub}</small>
    </column>
  </grid>
  <row class="flex row align--center">
    <column class="flex column w--100">
      <h6>{temp_mid} <sup>°</sup>C</h6>
    </column>

    <row
      class="flex row align--center gap--4 px--6 py--3 bg-primary theme-radius"
    >
      {#if temp_mid >= temp_sub && temp_mid <= temp_sup}
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
</a>
