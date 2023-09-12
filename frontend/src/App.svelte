<script lang="ts">
  import "ninja-keys";
  import CodeMirror from "svelte-codemirror-editor";
  import { Prec } from "@codemirror/state";
  import { keymap } from "@codemirror/view";
  import { ListManifests } from "../wailsjs/go/main/App";
  import CommandPalette from "./CommandPalette.svelte";

  let value = "";
  const manifestsP = ListManifests();
  let extensions = [
    Prec.highest(
      keymap.of([
        {
          key: "Mod-k",
          run: () => {
            const ninja = document.querySelector("ninja-keys");
            ninja.open();
            return true;
          },
        },
      ]),
    ),
  ];
</script>

<main>
  {#await manifestsP}
    <p>Loading...</p>
  {:then manifests}
    <CommandPalette {manifests}>
      <CodeMirror
        bind:value
        styles={{
          "&": {
            width: "100vw",
            height: "100vh",
          },
        }}
        {extensions}
      />
    </CommandPalette>
  {:catch error}
    <p>{error.message}</p>
  {/await}
</main>

<style>
</style>
