<script lang="ts">
  import { currentUser } from "$lib/stores/user";
  export let size: number;
</script>

<icon class="icon size--{size} circle bg-muted bord-primary bord--solid">
  <img src={$currentUser.Avatar} alt="avatar" loading="lazy" />
</icon>
