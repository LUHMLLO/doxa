<script lang="ts">
  const formData = {
    Email: "",
  };
</script>

<form on:submit|preventDefault class="flex column gap--24">
  <header class="w--100">
    <h4>Parece ser que olvidaste tu contraseña</h4>
    <p>
      No te preocupes, te enviaremos un correo con los pasos para resetear tu
      contraseña
    </p>
  </header>

  <fieldgroup class="w--100 flex column gap--16">
    <fieldset>
      <small>Correo electronico</small>
      <input type="text" bind:value={formData.Email} />
    </fieldset>
  </fieldgroup>

  <button
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    >Enviar Correo</button
  >
</form>
