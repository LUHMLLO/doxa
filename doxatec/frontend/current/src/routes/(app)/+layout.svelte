<script lang="ts">
  import Splashscreen from "$lib/components/splashscreen.svelte";
  import BaseLayout from "$lib/layouts/base.svelte";
  import { currentUser } from "$lib/stores";
  import { onMount } from "svelte";

  onMount(async () => {
    if (!$currentUser) {
      const res = await fetch("http://localhost:3000/api/auth/signature", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      const data = await res.json();
      currentUser.set(data);
    }
  });
</script>

{#if $currentUser}
  <BaseLayout>
    <slot />
  </BaseLayout>
{:else}
  <Splashscreen />
{/if}
