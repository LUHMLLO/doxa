<script lang="ts">
  import { goto } from "$app/navigation";
  import { currentUser } from "$lib/stores";
  import Avatar from "$lib/components/avatar.svelte";
  import Block_Section from "$lib/layouts/blocks/section.svelte";
  import Block_Header from "$lib/layouts/blocks/header.svelte";
</script>

<Block_Header>
  <banner
    class="block bg-accent theme-radius w--100 h--192p mx--auto offb--32"
  />
  <container class="container block">
    <row class="flex column md:row gap--8 w--100">
      <Avatar size={100} />
      <row class="flex row align--center w--100 md:mt--32">
        <column class="block w--100">
          <h5>{$currentUser?.username}</h5>
          <small>{$currentUser?.name}</small>
        </column>
        <button class="p--0 bord--hidden">
          <icon class="icon size--24">
            <i class="iconoir-lock" />
          </icon>
        </button>
      </row>
    </row>
  </container>
</Block_Header>

<Block_Section
  header="Plan Activo &amp; Metodo de pago"
  paragraph="Administra tu plan activo y el metodo de pago"
  container={"flex column gap--16"}
>
  <grid class="grid md:grid-cols--2 gap--16">
    <column class="flex column p--16 bg-light theme-radius">
      <h6>Plan Basico</h6>
      <small>Incluye: 1 Dispositivo</small>
      <hr />
      <footer class="flex row align--center">
        <button class="bg-muted clr-gray--2 bord--hidden theme-radius"
          >Actualizar plan</button
        >
      </footer>
    </column>
  </grid>
</Block_Section>

<Block_Section
  header="Cerrar sesión"
  paragraph="O cambiar de cuenta"
  container={"flex column gap--16"}
>
  <button
    on:click={() => goto("/signout")}
    class="bg-secondary clr-primary bord--hidden theme-radius"
    >Cerrar sesión
  </button>
</Block_Section>
