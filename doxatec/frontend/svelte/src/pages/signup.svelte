<script lang="ts">
  import { signUp } from "../lib/utils/signup";

  const user = {
    name: "",
    pass: "",
    fullname: "",
  };
</script>

<grid class="grid md:grid-cols--2 h-min--100vh">
  <column
    class="none md:flex justify--center align--center zindex--1 bg-accent h--100"
  >
    <container class="container block">
      <h3 class="text--center">Bienvenido a Doxatec</h3>
    </container>
  </column>

  <column class="flex justify--center align--center dpt-secondary zindex--2">
    <section>
      <container class="container block ">
        <form on:submit|preventDefault class="flex column gap--24">
          <header>
            <h4>Crea una cuenta para continuar</h4>
            <p>
              Ya tienes una cuenta?
              <a href="/auth/login">Inicia sesion</a>
            </p>
          </header>

          <fieldgroup class="flex column gap--16">
            <fieldset>
              <small>Nombre de usuario</small>
              <input type="text" bind:value={user.name} />
            </fieldset>
            <fieldset>
              <small>Nombre completo</small>
              <input type="text" bind:value={user.fullname} />
            </fieldset>
            <fieldset>
              <small>Contraseña</small>
              <input type="password" bind:value={user.pass} />
            </fieldset>
          </fieldgroup>

          <button
            on:click={() => signUp(user.name, user.pass, user.fullname)}
            class="bg-secondary clr-primary bord--hidden text--center"
            >Crear usuario</button
          >
        </form>
      </container>
    </section>
  </column>
</grid>
