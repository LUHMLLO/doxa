<script lang="ts">
  import { goto } from "$app/navigation";

  const formData = {
    username: "",
    password: "",
  };

  const RequestSignin = async () => {
    const res = await fetch("http://localhost:3000/api/auth/signin", {
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

    const data = await res.json();
    console.log(data);

    await goto("/");
  };
</script>

<form on:submit|preventDefault={RequestSignin} class="flex column gap--24">
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
      <input type="text" name="username" bind:value={formData.username} />
    </fieldset>
    <fieldset>
      <small>Contraseña</small>
      <input type="password" name="password" bind:value={formData.password} />
      <a href="/forgot" class="block w--100 text--10 text--end"
        >Olvidaste tu contraseña?</a
      >
    </fieldset>
  </fieldgroup>

  <button
    class="bg-secondary clr-primary bord--hidden text--center theme-radius"
    type="submit">Iniciar sesión</button
  >
</form>
