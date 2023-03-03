<script lang="ts">
  import type { PageData } from "./$types";
  import Device from "$lib/components/device.svelte";

  export let data: PageData;
</script>

<grid class="grid md:grid-cols--2 gap--16">
  {#each data.devices as d}
    <Device
      id={d.id}
      name={d.name}
      temp_sup={d.temp_sup}
      temp_mid={d.temp_mid}
      temp_sub={d.temp_sub}
    />
  {/each}
</grid>
