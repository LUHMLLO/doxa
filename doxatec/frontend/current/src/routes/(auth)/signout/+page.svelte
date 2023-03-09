<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";

    onMount(async () => {
        await fetch("http://localhost:3000/api/auth/signout", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
        });

        console.log("logged out");

        goto("/signin");
    });
</script>

<h1>logout</h1>
