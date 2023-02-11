<script lang="ts">
  let name: string;
  let category: string;
  let temp: number;

  export { name, category, temp };
</script>

<widget class="flex column bg-accent radius-theme p--16 dpt-dark">
  <h6 class="w--100">{name}</h6>
  <small class="weight--medium">{category}</small>

  <row class="flex row align--center">
    <p class="text--64 w--100">{temp}</p>

    <icon class="icon">
      {#if temp >= 20 || temp <= 10}
        <i class="iconoir-warning-triangle" />
      {:else}
        <i class="iconoir-lens" />
      {/if}
    </icon>
  </row>
</widget>
