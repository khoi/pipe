// import { EditorView } from '@codemirror/view';
// import { EditorState } from '@codemirror/state';
// import { basicSetup, minimalSetup } from '@uiw/codemirror-extensions-basic-setup';

import { langs } from "@uiw/codemirror-extensions-langs";
import { gruvboxDark, gruvboxLight } from "@uiw/codemirror-theme-gruvbox-dark";
import CodeMirror, { ReactCodeMirrorRef } from "@uiw/react-codemirror";
import React from "react";

import styles from "./App.module.css";
import { ListManifests, RunManifest } from "../wailsjs/go/main/App";
import { CommandPalette } from "./CommandPalette";
import "./global.css";
import { manifest } from "wailsjs/go/models";
import { Output } from "./types";
import { write } from "./output";
import useSystemTheme from "./useSystemTheme";

const extensions = [
  langs.css(),
  langs.javascript({ jsx: true }),
  langs.typescript(),
  langs.html(),
  langs.json(),
  langs.markdown(),
];

const emptyManifests: manifest.Manifest[] = [];

function App() {
  const valueRef = React.useRef<string>("");
  const codeMirrorRef = React.useRef<ReactCodeMirrorRef>(null);
  const setValue = React.useCallback((value: string) => {
    valueRef.current = value;
  }, []);
  const theme = useSystemTheme();

  const [manifests, setManifests] = React.useState(emptyManifests);
  const runManifest = React.useCallback(async (manifest: manifest.Manifest) => {
    try {
      const output = await RunManifest(manifest.id, valueRef.current);
      if (write(output, manifest.output as Output) === false) {
        return;
      }
      if (!codeMirrorRef.current || !codeMirrorRef.current.view) {
        return;
      }
      const view = codeMirrorRef.current.view;
      view.dispatch(
        view.state.changeByRange((range) => ({
          changes: [{ from: 0, to: view.state.doc.length, insert: output }],
          range: range,
        }))
      );
      view.focus();
    } catch (e) {
      console.error("Error running manifest", e);
    }
  }, []);

  React.useEffect(() => {
    async function loadManifests() {
      const manifests = await ListManifests();
      setManifests(manifests);
    }
    loadManifests();
  }, []);

  return (
    <React.Fragment>
      <CodeMirror
        ref={codeMirrorRef}
        className={styles.editor}
        value={valueRef.current}
        extensions={extensions}
        theme={theme == "dark" ? gruvboxDark : gruvboxLight}
        onChange={setValue}
      />
      <div className={styles.statusBar}>
        <div className={styles.statusBarItem}>Ln 1, Col 1</div>
      </div>
      <CommandPalette manifests={manifests} runManifest={runManifest} />
    </React.Fragment>
  );
}

export default App;
