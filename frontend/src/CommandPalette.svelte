<script lang="ts">
  import type { INinjaAction } from "ninja-keys/dist/interfaces/ininja-action";
  import { onMount } from "svelte";

  export let hotkeys: INinjaAction[] = [];

  onMount(async () => {
    const ninja = document.querySelector("ninja-keys");
    ninja.data = hotkeys;
  });
</script>

<svelte:head>
  <link
    href="https://fonts.googleapis.com/css?family=Material+Icons&display=swap"
    rel="stylesheet"
  />
</svelte:head>
<slot />
<ninja-keys placeholder="Run pipeline" />
