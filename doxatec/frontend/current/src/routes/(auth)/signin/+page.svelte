<script lang="ts">
  import { goto } from "$app/navigation";

  const formData = {
    username: "",
    password: "",
  };

  const RequestSignin = async () => {
    let response;

    try {
      response = await fetch("http://172.17.0.1:3000/api/auth/signin", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(formData),
      });
    } catch (error) {
      console.log("There was an error", error);
    }

    if (response?.ok) {
      const data = await response.json();
      console.log("response: ", data.message);
      goto("/");
    } else {
      response?.text().then((data) => {
        console.log("error: ", data);
      });
    }
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
