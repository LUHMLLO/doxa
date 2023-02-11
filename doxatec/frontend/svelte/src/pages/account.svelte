<script lang="ts">
  import { currentUser } from "../lib/stores/user";
  import { signOut } from "../lib/utils/signout";
</script>

{#if $currentUser}
  <header>
    <div class="m--0 bg-accent h--300p" />
    <container class="container block offt--56">
      <figure class="size--100 circle bg-muted bord-primary bord--solid" />
      <h4>{$currentUser.username}</h4>
      <p>{$currentUser.name}</p>
    </container>
  </header>
  <section>
    <container class="container block">
      <button class="bg-secondary clr-primary" on:click={() => signOut()}
        >Cerrar sesion
      </button>
    </container>
  </section>
{/if}
