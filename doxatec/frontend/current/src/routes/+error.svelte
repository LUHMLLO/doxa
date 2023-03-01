<script>
    import { page } from "$app/stores";
</script>

<grid class="grid h-min--100vh">
    <column class="none md:block zindex--1 bg-accent h--100" />
    <column class="flex justify--center align--center dpt-secondary zindex--2">
        <section class="w--100">
            <container class="container block">
                <h4>Parece ser que hubo un error con tu busqueda</h4>
                {#if $page?.error}
                    {#if $page?.status}
                        <p>Page status: {$page?.status}</p>
                    {/if}
                    {#if $page?.error?.message}
                        <p>Error message: {$page?.error?.message}</p>
                    {/if}
                {/if}
            </container>
        </section>
    </column>
    <column class="none md:block zindex--1 bg-accent h--100" />
</grid>
