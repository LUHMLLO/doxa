<script lang="ts">
  import { signUp } from "../lib/utils/signup";

  const user = {
    name: "",
    pass: "",
    fullname: "",
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
            <h4>Crea una cuenta para continuar</h4>
            <p>
              Ya tienes una cuenta?
              <a href="/auth/login" class="clr-tertiary">Inicia sesion</a>
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
            class="bg-dark clr-primary bord--hidden text--center"
            >Crear usuario</button
          >
        </form>
      </container>
    </section>
  </column>
</grid>
