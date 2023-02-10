<script lang="ts">
  import { logIn } from "../lib/utils/login";

  const user = {
    name: "",
    pass: "",
  };
</script>

<grid class="grid grid-cols--1 md:grid-cols--2 h--100vh">
  <column>
    <figure class="w--100 h--100">
      <img
        src="https://images.unsplash.com/photo-1558591710-4b4a1ae0f04d?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=774&q=80"
        alt="cover"
        class="w--100 h--100"
        loading="lazy"
      />
    </figure>
  </column>

  <column class="flex justify--center align--center">
    <section>
      <container class="container block ">
        <form on:submit|preventDefault class="flex column gap--24">
          <header>
            <h4>Bienvenido a Doxatec, inicie sesion para continuar</h4>
            <p>
              Aun no tienes una cuenta?
              <a href="/auth/login" class="clr-tertiary">Crear cuenta</a>
            </p>
          </header>

          <fieldgroup class="flex column gap--16">
            <fieldset>
              <small>Nombre de usuario</small>
              <input type="text" bind:value={user.name} />
            </fieldset>
            <fieldset>
              <small>Contraseña</small>
              <input type="password" bind:value={user.pass} />
              <a
                href="/auth/login"
                class="block clr-tertiary text--10 text--end w--100"
                >Olvidaste tu contraseña?</a
              >
            </fieldset>
          </fieldgroup>

          <button
            on:click={() => logIn(user.name, user.pass)}
            class="bg-dark clr-primary bord--hidden text--center"
            >Iniciar sesion</button
          >
        </form>
      </container>
    </section>
  </column>
</grid>
