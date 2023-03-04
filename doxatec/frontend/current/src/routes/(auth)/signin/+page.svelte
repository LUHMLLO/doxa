<script lang="ts">
  import { enhance } from '$app/forms'
</script>


<form method="POST" action="?/signin" use:enhance class="flex column gap--24">
  <header class="w--100">
    <h4>Bienvenido a Doxatec, inicie sesión para continuar</h4>
    <p>
      Aun no tienes una cuenta?
      <a href="/signup">Crear cuenta</a>
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
      <a href="/forgot" class="block w--100 text--10 text--end"
        >Olvidaste tu contraseña?</a
      >
    </fieldset>
  </fieldgroup>

  <button
    type="submit"
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    >Iniciar sesión</button
  >
</form>
