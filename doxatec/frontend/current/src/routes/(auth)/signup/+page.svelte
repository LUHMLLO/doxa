<script lang="ts">
  import { goto } from "$app/navigation";

  const formData = {
    username: "",
    password: "",
  };

  const CreateUser = async () => {
    await fetch("http://localhost:3000/api/auth/signin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        username: formData.username,
        password: formData.password,
      }),
    });

    await goto("/signin");
  };
</script>

<form on:submit|preventDefault={CreateUser} class="flex column gap--24">
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
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    type="submit">Crear usuario</button
  >
</form>
