<script lang="ts">
  import Navbar from "../components/navbar.svelte";
</script>

<wrapper
  class="overflow-hidden container block bg-primary clr-secondary dpt-dark radius-theme w--100vw"
>
  <main>
    <slot />
  </main>
</wrapper>
