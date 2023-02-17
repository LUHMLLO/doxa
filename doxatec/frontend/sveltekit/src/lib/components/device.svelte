<script lang="ts">
  export let name: string;
  export let temp: number;
</script>

<a href="/devices/001" class="bg-light clr-secondary flex w--100 column p--16 theme-radius">
  <grid class="grid grid-cols--2">
    <column class="block">
      <icon class="icon"><i class="iconoir-fridge" /></icon>
      <p>{name}</p>
    </column>
    <column class="block text--end">
      <small>Sup: 8</small>
      <small>Mid: 4</small>
      <small>Sub: 2</small>
    </column>
  </grid>
  <row class="flex row align--center">
    <column class="flex column w--100">
      <h6>{temp} <sup>°</sup>C</h6>
    </column>

    <row
      class="flex row align--center gap--4 px--6 py--3 bg-primary theme-radius"
    >
      {#if temp >= 2 && temp <= 8}
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
