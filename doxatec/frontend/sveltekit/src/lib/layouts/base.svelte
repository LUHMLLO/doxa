<script lang="ts">
  import Navbar from "$lib/components/navbar.svelte";
  import Sidebar from "$lib/components/sidebar.svelte";
</script>

<app
  class="flex bg-primary clr-secondary p--16 gap--16 w--100vw h--100vh ps--fixed inset"
>
  <Sidebar />
  <main class="h--100">
    <Navbar />
    <content class="block w--100 mt--16">
      <slot />
    </content>
  </main>
</app>
