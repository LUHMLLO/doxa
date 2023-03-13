<script lang="ts">
  import { currentUser } from "$lib/stores";

  export let size: number;
</script>

<icon
  class="icon size--{size} circle bg-muted bord-primary bord--solid bord-width--2"
>
  <img src={$currentUser?.avatar} alt="avatar" loading="lazy" />
</icon>
