<script lang="ts">
  import { signUp } from "./pocketbase";

  const user = {
    name: "",
    pass: "",
    fullname: "",
  };
</script>

<main>
  <section>
    <container class="container block">
      <form
        on:submit|preventDefault
        class="flex align--center column gap--64 w--300p m--auto"
      >
        <header class="w--100 text--center">
          <h1>Doxatec</h1>
          <p>Registra tu nueva cuenta aqui</p>
        </header>

        <fieldgroup class="w--100 flex column gap--8">
          <fieldset>
            <label for="">Nombre de usuario</label>
            <input type="text" bind:value={user.name} class="bord--hidden" />
            <hr class="m--0" />
          </fieldset>
          <fieldset>
            <label for="">Nombre completo</label>
            <input
              type="text"
              bind:value={user.fullname}
              class="bord--hidden"
            />
            <hr class="m--0" />
          </fieldset>
          <fieldset>
            <label for="">Contraseña</label>
            <input
              type="password"
              bind:value={user.pass}
              class="bord--hidden"
            />
            <hr class="m--0" />
          </fieldset>
        </fieldgroup>
        <footer class="w--100 text--center flex column align--center gap--8">
          <button
            on:click={() => signUp(user.name, user.pass, user.fullname)}
            class="w--100 bg-secondary clr-complementary radius-theme bord--hidden"
            >Crear usuario</button
          >
          <p>
            Ya tienes una cuenta?
            <a href="/auth/login" class="clr-tertiary">Inicia sesion</a>
          </p>
        </footer>
      </form>
    </container>
  </section>
</main>
