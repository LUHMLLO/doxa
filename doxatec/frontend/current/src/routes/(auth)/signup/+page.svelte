<script lang="ts">
  import { enhance } from "$app/forms";
</script>

<form method="POST" action="?/signup" use:enhance class="flex column gap--24">
  <header class="w--100">
    <h4>Crea una cuenta para continuar</h4>
    <p>
      Ya tienes una cuenta?
      <a href="/signin">Inicia sesión</a>
    </p>
  </header>

  <fieldgroup class="w--100 flex column gap--16">
    <fieldset>
      <small>Nombre de usuario</small>
      <input type="text" name="username" />
    </fieldset>
    <fieldset>
      <small>Contraseña</small>
      <input type="password" name="password" />
    </fieldset>
  </fieldgroup>

  <button
    type="submit"
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    >Crear usuario</button
  >
</form>
