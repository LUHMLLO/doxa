<script lang="ts">
  import Device from "$lib/components/device.svelte";
  import type { PageData } from "./$types";

  export let data: PageData;
</script>

<Device
  id={data.device.id}
  name={data.device.name}
  temp_sup={data.device.temp_sup}
  temp_mid={data.device.temp_mid}
  temp_sub={data.device.temp_sup}
/>
