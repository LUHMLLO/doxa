<script lang="ts">
  import { currentUser, signOut } from "./pocketbase";
</script>

<main>
  <section>
    <container class="container block">
      <header class="w--100 text--center">
        <h1>Bienvenido {$currentUser}</h1>
      </header>

      <button
        on:click={() => signOut()}
        class="w--100 bg-secondary clr-complementary radius-theme bord--hidden"
        >Cerrar sesion</button
      >
    </container>
  </section>
</main>
