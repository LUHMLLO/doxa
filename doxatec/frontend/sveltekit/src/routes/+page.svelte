<script lang="ts">
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  onMount(() => {
    setTimeout(() => {
      goto("/dashboard");
    }, 500);
  });
</script>

<div
  class="ps--fixed inset flex column justify--center align--center text--center w--100vw h--100vh"
>
  <h1>Doxatec</h1>
  <p>Cargando...</p>
</div>
