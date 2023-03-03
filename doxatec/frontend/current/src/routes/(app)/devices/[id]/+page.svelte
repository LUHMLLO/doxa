<script lang="ts">
  import Device from "$lib/components/device.svelte";
  import type { PageData } from "./$types";

  export let data: PageData;
</script>

<Device
  id={data.device.ID}
  owner={data.device.Owner}
  name={data.device.Name}
  temp_sup={data.device.TempSup}
  temp_mid={data.device.TempMid}
  temp_sub={data.device.TempSub}
/>
