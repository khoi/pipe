<script lang="ts">
  import CodeMirror from "svelte-codemirror-editor";
  import { Prec } from "@codemirror/state";
  import { keymap } from "@codemirror/view";
  import CommandPalette from "./CommandPalette.svelte";

  let value = "";
  let extensions = [
    Prec.highest(
      keymap.of([
        {
          key: "Mod-k",
          run: () => {
            console.log("hi");
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
  <CommandPalette>
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
</main>

<style>
</style>
