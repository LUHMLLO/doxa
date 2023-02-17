<script lang="ts">
  import Avatar from "$lib/components/avatar.svelte";
  import { currentUser } from "$lib/stores/user";
  import { goto } from "$app/navigation";
</script>

<header>
  <div class="m--0 bg-accent h--192p theme-radius" />
  <container class="container block offt--32">
    <row class="flex row align--center gap--8 w--100">
      <Avatar size={100} />
      <row class="flex row align--center w--100 mt--32">
        <column class="flex column w--100">
          <h4>{$currentUser.Username}</h4>
          <p>{$currentUser.Customer.Name}</p>
        </column>
        <button class="p--0 bord--hidden">
          <icon class="icon size--24">
            <i class="iconoir-lock" />
          </icon>
        </button>
      </row>
    </row>
  </container>
</header>

<container class="container flex column gap--16">
  <header>
    <h6>Plan Activo &amp; Metodo de pago</h6>
    <p>Administra tu plan activo y el metodo de pago</p>
  </header>
  <grid class="grid md:grid-cols--2 gap--16">
    <column class="flex column p--16 bg-light theme-radius">
      <h6>Plan Basico</h6>
      <small>Incluye: 1 Dispositivo</small>
      <hr />
      <footer class="flex row align--center">
        <button class="bg-secondary clr-primary bord-hidden theme-radius"
          >Actualizar plan</button
        >
      </footer>
    </column>
  </grid>
</container>

<container class="container flex column gap--16">
  <header>
    <h6>Cerrar sesión</h6>
    <p>O cambiar de cuenta</p>
  </header>
  <button
    class="bg-secondary clr-primary bord-hidden theme-radius"
    on:click={() => goto("/login")}
    >Cerrar sesión
  </button>
</container>
