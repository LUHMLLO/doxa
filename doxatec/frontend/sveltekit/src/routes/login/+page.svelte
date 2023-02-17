<script lang="ts">
  import type User from "$lib/types/user";

  const user: User = {};
</script>

<form on:submit|preventDefault class="flex column gap--24">
  <header class="w--100">
    <h4>Bienvenido a Doxatec, inicie sesion para continuar</h4>
    <p>
      Aun no tienes una cuenta?
      <a href="/signup">Crear cuenta</a>
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
      <a href="/forgot" class="block text--10 text--end"
        >Olvidaste tu contraseña?</a
      >
    </fieldset>
  </fieldgroup>

  <button
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    >Iniciar sesion</button
  >
</form>
