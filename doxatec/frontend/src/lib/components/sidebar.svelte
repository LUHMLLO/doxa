<script lang="ts">
  import { Link } from "svelte-navigator";

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
    {
      icon: "iconoir-user-circle",
      name: "my account",
      path: "/account",
    },
  ];
</script>

<nav
  class="column align--center gap--8 px--8 py--16 bg-light bord-primary bord--solid bord-width--2 radius-theme h--100"
>
  <Link to="/" class="inline-flex justify--center align--center mb--16">
    <icon
      class="icon size--48 circle my--auto bord--solid bord-width--1 bord-primary"
    >
      <img src="logo.svg" alt="logo" class="" />
    </icon>
  </Link>

  {#each buttons as button}
    <Link
      to={button.path}
      class="inline-flex justify--center align--center p--8 bg-primary clr-secondary radius-theme"
    >
      <icon class="icon size--24">
        <i class={button.icon} />
      </icon>
    </Link>
  {/each}
</nav>
