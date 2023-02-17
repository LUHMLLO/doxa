<script lang="ts">
  import Dash from "$lib/components/dash.svelte";
  import Navbar from "$lib/components/navbar.svelte";
  import Sidebar from "$lib/components/sidebar.svelte";
</script>

<app
  class="flex bg-primary clr-secondary w--100vw h--100vh ps--fixed inset overflow--hidden"
>
  <Sidebar />
  <main class="flex column sm:my--16">
    <Navbar />
    <content class="block h--100 shrink py--16 mx--16 scroll--y">
      <slot />
    </content>
    <Dash />
  </main>
</app>
