<script lang="ts">
  import { goto } from "$app/navigation";
  import { currentUser } from "$lib/stores/user";
</script>

<form on:submit|preventDefault class="flex column gap--24">
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
      <input type="text" bind:value={$currentUser.Username} />
    </fieldset>
    <fieldset>
      <small>Contraseña</small>
      <input type="password" bind:value={$currentUser.Password} />
      <a href="/forgot" class="block w--100 text--10 text--end"
        >Olvidaste tu contraseña?</a
      >
    </fieldset>
  </fieldgroup>

  <button
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    on:click={() => goto("/dashboard")}>Iniciar sesión</button
  >
</form>
