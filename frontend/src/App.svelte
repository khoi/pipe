<script lang="ts">
  import CodeMirror from "svelte-codemirror-editor";
  import CommandPalette from "./CommandPalette.svelte";

  let value = "";
</script>

<main>
  <CommandPalette />
  <CodeMirror
    bind:value
    styles={{
      "&": {
        width: "100vw",
        height: "100vh",
      },
    }}
  />
</main>

<style>
</style>
