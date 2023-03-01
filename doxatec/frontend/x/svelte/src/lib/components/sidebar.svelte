<script lang="ts">
  import { Link } from "svelte-navigator";
  import Avatar from "../atoms/avatar.svelte"

  const buttons = [
    {
      icon: "iconoir-report-columns",
      name: "dashboard",
      path: "/",
    },
    {
      icon: "iconoir-fridge",
      name: "my devices",
      path: "/devices",
    },
  ];
</script>

<nav
  class="column align--center gap--16 px--8 py--16 bg-light bord-primary bord--solid bord-width--2 radius-theme h--100 dpt-secondary"
>
  <Link to="/" class="inline-flex justify--center align--center">
    <icon
      class="icon size--48 circle my--auto bord--solid bord-width--1 bord-primary"
    >
      <img src="logo.svg" alt="logo" class="" />
    </icon>
  </Link>

  <nav class="column grow">
    {#each buttons as button}
      <Link
        to={button.path}
        class="inline-flex justify--center align--center p--8 clr-secondary radius-theme"
      >
        <icon class="icon size--24">
          <i class={button.icon} />
        </icon>
      </Link>
    {/each}
  </nav>

  <icon class="icon size--24">
    <i class="iconoir-settings" />
  </icon>
  <Avatar size="size--48"/>
</nav>
