<script lang="ts">
  import Navbar from "../components/navbar.svelte";
  import Sidebar from "../components/sidebar.svelte";
</script>

<app class="block bg-primary clr-secondary w--100vw h--100vh ps--fixed inset">
  <Navbar />
  <row class="flex row">
    <Sidebar />
    <main class="w--100">
      <slot />
    </main>
  </row>
</app>
