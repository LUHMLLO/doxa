<script lang="ts">
  import Avatar from "$lib/components/avatar.svelte";
  const buttons = [
    {
      icon: "iconoir-report-columns",
      name: "dashboard",
    },
    {
      icon: "iconoir-computer",
      name: "devices",
    },
  ];
</script>

<nav
  class="sm:none justify--space-evenly align--center gap--16 h--64p px--12 py--6 bg-primary bord-light bord--solid bord-width--2 theme-radius zindex--12 dpt-secondary"
>
  {#each buttons as button}
    <a
      href="/{button.name}"
      class="clr-secondary inline-flex justify--center align--center mx--auto"
    >
      <icon class="icon size--32">
        <i class={button.icon} />
      </icon>
    </a>
  {/each}

  <a
    href="/settings"
    class="clr-secondary inline-flex justify--center align--center mx--auto"
  >
    <icon class="icon size--32">
      <i class="iconoir-settings" />
    </icon>
  </a>
  <a
    href="/profile"
    class="clr-secondary inline-flex justify--center align--center mx--auto"
  >
    <Avatar size={32} />
  </a>
</nav>
