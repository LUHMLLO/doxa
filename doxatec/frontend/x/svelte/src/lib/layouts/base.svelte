<script lang="ts">
  import Navbar from "../components/navbar.svelte";
  import Sidebar from "../components/sidebar.svelte";
</script>

<app class="flex gap--8 bg-primary clr-secondary w--100vw h--100vh ps--fixed inset">
  <Sidebar />
  <main class="w--100 h--100vh overflow-hidden scroll-y py--8">
    <slot />
  </main>
</app>
