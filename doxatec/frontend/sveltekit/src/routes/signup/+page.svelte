<script lang="ts">
  import { goto } from "$app/navigation";
  import { currentUser } from "$lib/stores/user";
</script>

<form on:submit|preventDefault class="flex column gap--24">
  <header class="w--100">
    <h4>Crea una cuenta para continuar</h4>
    <p>
      Ya tienes una cuenta?
      <a href="/login">Inicia sesión</a>
    </p>
  </header>

  <fieldgroup class="w--100 flex column gap--16">
    <fieldset>
      <small>Nombre de usuario</small>
      <input type="text" bind:value={$currentUser.Username} />
    </fieldset>
    <fieldset>
      <small>Contraseña</small>
      <input type="password" bind:value={$currentUser.Password} />
    </fieldset>
  </fieldgroup>

  <button
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    on:click={() => goto("/dashboard")}>Crear usuario</button
  >
</form>
