<script lang="ts">
  import { currentUser } from "../stores/user";

  export let size: string = 'size--100';
</script>

<icon class="icon {size} circle bg-muted bord-primary bord--solid">
  <img src={$currentUser.avatar || 'https://cdn.dribbble.com/userupload/2798814/file/original-3cfdbabadfd8f92aed97b0c0b57c6b89.png?compress=1&resize=752x'} alt="avatar" />
</icon>
