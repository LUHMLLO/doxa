<script lang="ts">
    export let section: string = "";
    export let container: string = "";

    export let header: string = "";
    export let paragraph: string = "";
</script>

<section class="section {section}">
    <container class="container {container}">
        {#if header || paragraph}
            <header>
                {#if header}
                    <h6>{header}</h6>
                {/if}
                {#if paragraph}
                    <p>{paragraph}</p>
                {/if}
            </header>
        {/if}
        <slot />
    </container>
</section>
