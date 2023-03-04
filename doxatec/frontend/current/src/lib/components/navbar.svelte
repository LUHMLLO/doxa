<script>
  import { page } from "$app/stores";
</script>

<nav
  class="align--center gap--16 h--32p sm:mx--16 bg-primary bord-light bord--solid bord-width--2 theme-radius zindex--12 dpt-secondary"
>
  <a href={void 0} class="inline-flex justify--center align--center">
    <button class="flex row align--center gap--8 bg-light bord--hidden">
      <h6 class="text--capitalize w--100 whitespace--nowrap">
        {#if $page.url.pathname == "/"}
          {$page.url.pathname.replace("/", "Dashboard")}
        {/if}
        {$page.url.pathname.replace("/", "")}
      </h6>
    </button>
    <tooltip class="tooltip dpt-secondary">Ruta actual</tooltip>
  </a>

  <nav class="justify--end align--center gap--16 w--100">
    <button class="p--0 bord--hidden">
      <icon class="icon">
        <i class="iconoir-search" />
      </icon>
    </button>
    <button class="p--0 bord--hidden">
      <icon class="icon">
        <i class="iconoir-bell" />
      </icon>
    </button>
  </nav>

  <button class="flex row align--center gap--8 bg-light bord--hidden">
    <icon class="icon">
      <i class="iconoir-cloud-check" />
    </icon>
  </button>
</nav>
