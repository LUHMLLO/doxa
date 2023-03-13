<script lang="ts">
    import type { LayoutData } from "./$types";
    import { onMount } from "svelte";
    import { currentUser } from "$lib/stores";
    import Splashscreen from "$lib/components/splashscreen.svelte";
    import BaseLayout from "$lib/layouts/base.svelte";

    export let data: LayoutData;

    onMount(() => {
        currentUser.set(data.user);
    });
</script>

{#if $currentUser}
    <BaseLayout>
        <slot />
    </BaseLayout>
{:else}
    <Splashscreen />
{/if}
