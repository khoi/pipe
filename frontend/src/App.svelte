<script lang="ts">
  import "ninja-keys";
  import CodeMirror from "svelte-codemirror-editor";
  import { Prec } from "@codemirror/state";
  import { keymap } from "@codemirror/view";
  import { ListManifests, RunManifest } from "../wailsjs/go/main/App";
  import CommandPalette from "./CommandPalette.svelte";
  import { write } from "./output";
  import type { Output } from "./types";
  import { oneDark } from "@codemirror/theme-one-dark";

  let value = "";
  const hotkeysP = ListManifests().then((manifests) => {
    return manifests.map((manifest, idx) => {
      return {
        id: manifest.name,
        title: `${manifest.name} - ${manifest.description}`,
        hotkey: idx < 9 ? `cmd+${idx + 1}` : undefined,
        mdIcon: "apps",
        handler: async () => {
          console.log(`Running ${manifest.name} for input\n${value}`);
          try {
            const output = await RunManifest(manifest.id, value);
            if (write(output, manifest.output as Output) === false) {
              return;
            }
            value = output;
          } catch (e) {
            console.error("Error running manifest", e);
          }
        },
      };
    });
  });

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
      ])
    ),
  ];
</script>

<main>
  {#await hotkeysP}
    <p>Loading...</p>
  {:then hotkeys}
    <CommandPalette {hotkeys}>
      <CodeMirror
        basic={true}
        lineWrapping={true}
        bind:value
        styles={{
          "&": {
            width: "100vw",
            height: "100vh",
          },
        }}
        {extensions}
        theme={oneDark}
      />
    </CommandPalette>
  {:catch error}
    <p>{error.message}</p>
  {/await}
</main>

<style>
</style>
