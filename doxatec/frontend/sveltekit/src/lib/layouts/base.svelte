<script lang="ts">
  import Bottombar from "$lib/components/bottombar.svelte";
  import Navbar from "$lib/components/navbar.svelte";
  import Sidebar from "$lib/components/sidebar.svelte";
</script>

<app
  class="flex bg-primary clr-secondary w--100vw h--100vh ps--fixed inset overflow--hidden"
>
  <Sidebar />
  <main class="h--100 sm:mt--16 mx--16 scroll--y">
    <Navbar />
    <content class="block w--100 my--48 sm:my--0">
      <slot />
    </content>
  </main>
  <Bottombar />
</app>
