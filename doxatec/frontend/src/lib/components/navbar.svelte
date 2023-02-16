<script lang="ts">
  import { Link } from "svelte-navigator";
</script>

<nav class="justify--space-between align--center gap--32 px--24 py--12 bg-light">
  <Link to="/" class="inline-flex align--center">
    <icon class="icon size--48 circle my--auto">
      <img src="logo.svg" alt="logo" class="" />
    </icon>
  </Link>

  <button class="p--0 bord--hidden">
    <icon class="icon size--24 circle">
      <i class="iconoir-search" />
    </icon>
  </button>
</nav>
