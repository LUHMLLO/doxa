<script lang="ts">
  import { currentUser } from "../lib/stores/user";
  import { signOut } from "../lib/utils/signout";
</script>

{#if $currentUser}
  <section>
    <container class="container block">
      <header class="w--100 text--center">
        <h1>Bienvenido {$currentUser.username}</h1>
      </header>

      <button
        on:click={() => signOut()}
        class="w--100 bg-secondary clr-complementary radius-theme bord--hidden"
        >Cerrar sesion</button
      >
    </container>
  </section>
{/if}
