<form action="/login" method="POST" class="flex column gap--24">
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
      <input type="text" />
    </fieldset>
    <fieldset>
      <small>Contraseña</small>
      <input type="password" />
    </fieldset>
  </fieldgroup>

  <button
    type="submit"
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    >Crear usuario</button
  >
</form>
