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
  import { json } from "@codemirror/lang-json";

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
      ]),
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
            height: "calc(100vh - calc(var(--status-height) * 2))",
          },
        }}
        {extensions}
        theme={oneDark}
        lang={json()}
      />
    </CommandPalette>
    <div class="status">
      <span>Ctrl+K to open command palette</span>
    </div>
  {:catch error}
    <p>{error.message}</p>
  {/await}
</main>

<style>
  :root {
    --status-height: 20px;
  }
  .status {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    height: var(--status-height);
    padding: 0.5rem;
    color: #999;
    font-size: 0.75rem;
    display: flex;
  }
  .status span {
    display: flex;
    align-items: center;
  }
</style>
