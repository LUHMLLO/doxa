<script lang="ts">
  import { onMount, onDestroy } from "svelte";
  import pb from "../lib/utils/pocketbase";
  import { currentUser } from "../lib/stores/user";
  import { signOut } from "../lib/utils/signout";
  import DeviceTemperature from "../lib/widgets/deviceTemperature.svelte";

  let newDeviceName: string;
  let devices = [];
  let unsubscribe: () => void;

  async function createNewDevice() {
    const data = {
      name: newDeviceName,
      owner: $currentUser.id,
    };

    await pb.collection("devices").create(data);
  }

  onMount(async () => {
    const devicesList = await pb.collection("devices").getList(1, 20, {
      sort: "created",
      expand: "user",
    });
    devices = devicesList.items;

    unsubscribe = await pb
      .collection("devices")
      .subscribe("*", async ({ action, record }) => {
        if (action === "create") {
          devices = [...devices, record];
        }
        if (action === "delete") {
          devices = devices.filter((device) => device.id !== record.id);
        }
      });
  });

  onDestroy(() => {
    unsubscribe();
  });
</script>

{#if $currentUser}
  <header>
    <div class="m--0 bg-accent h--300p" />
    <container class="container block offt--56">
      <figure class="size--100 circle bg-muted bord-primary bord--solid" />
      <h4>{$currentUser.username}</h4>
      <p>{$currentUser.name}</p>
    </container>
  </header>
  <section>
    <container class="container block">
      <grid class="grid grid-cols--2 md:grid-cols--4 gap--16">
        {#each devices as device (device.id)}
          <DeviceTemperature
            name={device.name}
            category={device.category}
            temp={Math.floor(Math.random() * 24)}
          />
        {/each}
        <widget class="flex column bord-accent bord--dashed radius-theme p--16">
          <fieldgroup>
            <fieldset>
              <small>Device name</small>
              <input type="text" bind:value={newDeviceName} />
            </fieldset>
          </fieldgroup>
          <button class="w--100 bg-secondary clr-primary" on:click={() => createNewDevice()}
            >Add device
          </button>
        </widget>
      </grid>
    </container>
  </section>
  <section>
    <container class="container block">
      <button class="bg-secondary clr-primary" on:click={() => signOut()}
        >Cerrar sesion
      </button>
    </container>
  </section>
{/if}
