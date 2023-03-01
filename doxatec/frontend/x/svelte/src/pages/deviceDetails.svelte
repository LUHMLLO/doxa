<script lang="ts">
  export let layout: any;
  export let id: string;
</script>

<svelte:component this={layout}>
  <h2>ID: {id}</h2>
</svelte:component>
