<script lang="ts">
  import type User from "$lib/types/user";

  const user: User = {
    ID: "",
    Username: "",
    Password: "",
  };
</script>

<form on:submit|preventDefault class="flex column gap--24">
  <header class="w--100">
    <h4>Crea una cuenta para continuar</h4>
    <p>
      Ya tienes una cuenta?
      <a href="/login">Inicia sesion</a>
    </p>
  </header>

  <fieldgroup class="w--100 flex column gap--16">
    <fieldset>
      <small>Nombre de usuario</small>
      <input type="text" bind:value={user.Username} />
    </fieldset>
    <fieldset>
      <small>Contraseña</small>
      <input type="password" bind:value={user.Password} />
    </fieldset>
  </fieldgroup>

  <button
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    >Crear usuario</button
  >
</form>
