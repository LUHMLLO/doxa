<script lang="ts">
  import { logIn } from "./pocketbase";

  const user = {
    name: "",
    pass: "",
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
          <p>Inicie sesion en su cuenta</p>
        </header>

        <fieldgroup class="w--100 flex column gap--8">
          <fieldset>
            <label for="">Nombre de usuario</label>
            <input type="text" bind:value={user.name} class="bord--hidden" />
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
            on:click={() => logIn(user.name, user.pass)}
            class="w--100 bg-secondary clr-complementary radius-theme bord--hidden"
            >Iniciar sesion</button
          >
          <p>
            Aun no tienes una cuenta?
            <a href="/auth/register" class="clr-tertiary">Crear cuenta</a>
          </p>
        </footer>
      </form>
    </container>
  </section>
</main>
