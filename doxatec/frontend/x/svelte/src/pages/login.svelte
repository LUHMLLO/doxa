<script lang="ts">
  import { logIn } from "../lib/utils/login";

  const user = {
    name: "",
    pass: "",
  };
</script>

<grid class="grid md:grid-cols--2 h-min--100vh">
  <column class="flex justify--center align--center dpt-secondary zindex--2">
    <section>
      <container class="container block">
        <form on:submit|preventDefault class="flex column gap--24">
          <header>
            <h4>Bienvenido a Doxatec, inicie sesion para continuar</h4>
            <p>
              Aun no tienes una cuenta?
              <a href="/auth/login">Crear cuenta</a>
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
              <a href="/auth/login" class="block text--10 text--end"
                >Olvidaste tu contraseña?</a
              >
            </fieldset>
          </fieldgroup>

          <button
            on:click={() => logIn(user.name, user.pass)}
            class="bg-secondary clr-primary bord--hidden text--center"
            >Iniciar sesion</button
          >
        </form>
      </container>
    </section>
  </column>

  <column
    class="none md:flex justify--center align--center zindex--1 bg-accent h--100"
  >
    <container class="container block">
      <h3 class="text--center">Bienvenido a Doxatec</h3>
    </container>
  </column>
</grid>
