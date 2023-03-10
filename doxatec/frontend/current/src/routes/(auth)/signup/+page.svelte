<script lang="ts">
  import { goto } from "$app/navigation";

  const formData = {
    username: "",
    password: "",
    avatar:
      "https://cdn.dribbble.com/users/957405/screenshots/16190682/media/07e2dcd89bd318885a354426bb403d22.png",
    name: "",
    email: "",
    phone: "",
    role: "user",
  };

  const RequestSignup = async () => {
    await fetch("http://localhost:3000/api/auth/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(formData),
    });

    await goto("/signin");
  };
</script>

<form on:submit|preventDefault={RequestSignup} class="flex column gap--24">
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
      <input type="text" name="username" bind:value={formData.username} />
    </fieldset>
    <fieldset>
      <small>Contraseña</small>
      <input type="password" name="password" bind:value={formData.password} />
    </fieldset>
    <fieldset>
      <small>Nombre Completo</small>
      <input type="text" name="name" bind:value={formData.name} />
    </fieldset>
    <fieldset>
      <small>Email</small>
      <input type="email" name="email" bind:value={formData.email} />
    </fieldset>
    <fieldset>
      <small>Phone</small>
      <input type="tel" name="tel" bind:value={formData.phone} />
    </fieldset>
  </fieldgroup>

  <button
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    type="submit">Crear usuario</button
  >
</form>
