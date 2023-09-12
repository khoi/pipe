<script lang="ts">
  import type { INinjaAction } from "ninja-keys/dist/interfaces/ininja-action";
  import { onMount } from "svelte";
  import type { manifest } from "wailsjs/go/models";

  export let manifests: manifest.Manifest[] = [];

  const hotkeys: INinjaAction[] = manifests.map((manifest, idx) => {
    return {
      id: manifest.name,
      title: `${manifest.name} - ${manifest.description}`,
      hotkey: idx < 9 ? `cmd+${idx + 1}` : undefined,
      mdIcon: "apps",
      handler: () => {
        console.log(`navigation to ${manifest.name}`);
      },
    };
  });

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
