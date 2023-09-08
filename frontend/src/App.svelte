<script lang="ts">
   import CodeMirror from "svelte-codemirror-editor";

    let value = "";
</script>

<main>
  <CodeMirror bind:value styles={{
    "&": {
        width: "100vw",
        height: "100vh",
    },
}}/>
</main>

<style>
</style>
