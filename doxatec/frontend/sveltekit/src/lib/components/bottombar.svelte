<script lang="ts">
  import Avatar from "$lib/components/avatar.svelte";
  const buttons = [
    {
      icon: "iconoir-compass",
      name: "dashboard",
    },
    {
      icon: "iconoir-computer",
      name: "devices",
    },
  ];
</script>

<bottombar
  class="flex sm:none row justify--space-between align--center gap--32 px--16 py--8 w--100 bg-light bord-light bord--solid bord-width--2 theme-radius zindex--12 ps--fixed right bottom left"
>
  {#each buttons as button}
    <a
      href="/{button.name}"
      class="clr-secondary inline-flex justify--center align--center"
    >
      <icon class="icon size--24">
        <i class={button.icon} />
      </icon>
    </a>
  {/each}

  <a
    href="/settings"
    class="clr-secondary inline-flex justify--center align--center"
  >
    <icon class="icon size--24">
      <i class="iconoir-settings" />
    </icon>
  </a>
  <a
    href="/account"
    class="clr-secondary inline-flex justify--center align--center"
  >
    <Avatar size={24} />
  </a>
</bottombar>
