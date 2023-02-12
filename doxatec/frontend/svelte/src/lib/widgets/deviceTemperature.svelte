<script lang="ts">
  let name: string;
  let category: string;
  let temp: number;

  export { name, category, temp };
</script>

<widget class="flex column bg-accent radius-theme p--16 dpt-secondary">
  <h6 class="w--100">{name}</h6>
  <small class="weight--medium">{category}</small>

  <row class="flex row align--center">
    <p class="text--64">{temp}</p>

    <icon class="icon">
      {#if temp >= 2 || temp <= 8}
        <i class="iconoir-lens" />
      {:else}
        <i class="iconoir-warning-triangle" />
      {/if}
    </icon>
  </row>
</widget>
