<script lang="ts">
  import Device from "$lib/components/device.svelte";

  const Devices = [
    {
      name: "Device #01",
      temp: Math.floor(Math.random() * 12),
    },
    {
      name: "Device #02",
      temp: Math.floor(Math.random() * 24),
    },
    {
      name: "Device #03",
      temp: Math.floor(Math.random() * 8),
    },
    {
      name: "Device #04",
      temp: Math.floor(Math.random() * 16),
    },
    {
      name: "Device #05",
      temp: Math.floor(Math.random() * 32),
    },
  ];
</script>

<grid class="grid md:grid-cols--2 gap--16">
  {#each Devices as d}
    <Device name={d.name} temp={d.temp} />
  {/each}
</grid>
