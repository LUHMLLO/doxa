<script lang="ts">
  import Avatar from "$lib/components/avatar.svelte";
</script>

<sidebar
  class="none sm:flex column align--center gap--32 w--64p py--16 ml--16 my--16 bord-light bord--solid bord-width--2 theme-radius zindex--12 dpt-secondary bd-blur--8"
  style="background-color:rgba(240, 240, 240, 0.8);"
>
  <nav class="column gap--8">
    <a href="/" class="clr-secondary inline-flex justify--center align--center">
      <icon class="icon size--32">
        <img src="/logo.svg" alt="logo" loading="lazy" />
      </icon>
      <tooltip class="tooltip dpt-secondary">Volver al inicio</tooltip>
    </a>
  </nav>

  <nav class="column gap--8 align--center grow">
    <a
      href="/dashboard"
      class="clr-secondary inline-flex justify--center align--center"
    >
      <icon class="icon size--24">
        <i class="iconoir-compass" />
      </icon>
      <tooltip class="tooltip dpt-secondary">dashboard</tooltip>
    </a>
    <a
      href="/devices"
      class="clr-secondary inline-flex justify--center align--center"
    >
      <icon class="icon size--24">
        <i class="iconoir-computer" />
      </icon>
      <tooltip class="tooltip dpt-secondary">dispositivos</tooltip>
    </a>
  </nav>

  <nav class="column gap--8 align--center">
    <a
      href="/settings"
      class="clr-secondary inline-flex justify--center align--center"
    >
      <icon class="icon size--24">
        <i class="iconoir-settings" />
      </icon>
      <tooltip class="tooltip dpt-secondary">ajustes</tooltip>
    </a>
    <a
      href="/account"
      class="clr-secondary inline-flex justify--center align--center"
    >
      <Avatar size={24} />
      <tooltip class="tooltip dpt-secondary">mi cuenta</tooltip>
    </a>
  </nav>
</sidebar>
