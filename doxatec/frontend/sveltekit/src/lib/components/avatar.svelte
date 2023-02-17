<script lang="ts">
  import { currentUser } from "$lib/stores/user";
  export let size: number;
</script>

<icon
  class="icon size--{size} circle bg-muted bord-primary bord--solid bord-width--2"
>
  <img src={$currentUser.Avatar} alt="avatar" loading="lazy" />
</icon>
