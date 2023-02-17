<script lang="ts">
  import Navbar from "$lib/components/navbar.svelte";
  import Sidebar from "$lib/components/sidebar.svelte";
</script>

<app
  class="flex bg-primary clr-secondary w--100vw h--100vh ps--fixed inset overflow--hidden"
>
  <Sidebar />
  <main class="h--100 mt--16 mr--16 scroll--y">
    <Navbar />
    <content class="block w--100">
      <slot />
    </content>
  </main>
</app>
