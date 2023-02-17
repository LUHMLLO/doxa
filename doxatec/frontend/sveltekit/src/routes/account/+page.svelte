<script lang="ts">
  import Avatar from "$lib/components/avatar.svelte";
  import { currentUser } from "$lib/stores/user";
</script>

<header>
  <div class="m--0 bg-accent h--192p theme-radius" />
  <container class="container block offt--32">
    <row class="flex row align--center gap--8">
      <Avatar size={100} />
      <column class="flex column mt--32">
        <h4>{$currentUser.Username}</h4>
        <p>{$currentUser.Name}</p>
      </column>
    </row>
  </container>
</header>
<section>
  <container class="container block">
    <h6>Plan Activo &amp; Metodo de pago</h6>
    <p>Administra tu plan activo y el metodo de pago</p>
    <grid class="grid md:grid-cols--2 gap--16" />
  </container>
</section>
<section>
  <container class="container block">
    <button class="bg-secondary clr-primary">Cerrar sesion </button>
  </container>
</section>
