<script lang="ts">
  import type { PageData } from "./$types";
  import Device from "$lib/components/device.svelte";

  export let data: PageData;
</script>

<grid class="grid md:grid-cols--2 gap--16">
  {#each data.devices as d}
    <Device
      id={d.ID}
      owner={d.Owner}
      name={d.Name}
      temp_sup={d.TempSup}
      temp_mid={d.TempMid}
      temp_sub={d.TempSub}
    />
  {/each}
</grid>
