<script lang="ts">
  import { NinjaKeys } from "ninja-keys";
  import { onMount } from "svelte";
  const hotkeys = [
    {
      id: "Home",
      title: "Open Home",
      hotkey: "cmd+h",
      mdIcon: "home",
      handler: () => {
        console.log("navigation to home");
      },
    },
    {
      id: "Open Projects",
      title: "Open Projects",
      hotkey: "cmd+p",
      mdIcon: "apps",
      handler: () => {
        console.log("navigation to projects");
      },
    },
  ];

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
<h1>Ctrl+K</h1>
<ninja-keys placeholder="My custom placeholder" />
