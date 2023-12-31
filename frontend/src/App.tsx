import { githubLight, githubDark } from "@uiw/codemirror-theme-github";
import CodeMirror, {
  Compartment,
  ReactCodeMirrorRef,
} from "@uiw/react-codemirror";
import React from "react";

import styles from "./App.module.css";
import { ListManifests, RunManifest } from "../wailsjs/go/main/App";
import { CommandPalette } from "./CommandPalette";
import "./global.css";
import { manifest } from "wailsjs/go/models";
import useSystemTheme from "./useSystemTheme";
import { Loader2 } from "lucide-react";
import { write } from "./output";
import { Output } from "./types";
import * as detect from "./languageDetection/detect";

const languageConf = new Compartment();
const emptyManifests: manifest.Manifest[] = [];

function App() {
  const valueRef = React.useRef<string>("");
  const codeMirrorRef = React.useRef<ReactCodeMirrorRef>(null);
  const handleLangDetection = React.useCallback(
    (lang: detect.DetectionResult) => {
      if (!codeMirrorRef.current || !codeMirrorRef.current.view) {
        return;
      }

      console.log("Reconfiguring language", lang);

      codeMirrorRef.current.view.dispatch({
        effects: languageConf.reconfigure(lang),
      });
    },
    [],
  );

  const setValue = React.useCallback((value: string) => {
    valueRef.current = value;
    detect.detectLanguage(valueRef.current, handleLangDetection);
  }, []);
  const theme = useSystemTheme();

  const [manifests, setManifests] = React.useState(emptyManifests);
  React.useEffect(() => {
    async function loadManifests() {
      const manifests = await ListManifests();
      setManifests(manifests);
    }
    loadManifests();
  }, []);

  const [loading, setLoading] = React.useState(false);
  const runManifest = React.useCallback(async (manifest: manifest.Manifest) => {
    try {
      setLoading(true);
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
        })),
      );
      view.focus();
    } catch (error) {
      console.error("Error running manifest", error);
    } finally {
      setLoading(false);
    }
  }, []);

  return (
    <React.Fragment>
      <CodeMirror
        readOnly={loading}
        basicSetup={true}
        ref={codeMirrorRef}
        className={styles.editor}
        value={valueRef.current}
        extensions={[languageConf.of([])]}
        theme={theme == "dark" ? githubDark : githubLight}
        onChange={setValue}
      />
      <div className={styles.statusBar}>
        <div className={styles.statusBarItem}>
          Press <kbd>⌘</kbd> + <kbd>K</kbd> to open the command palette
        </div>
        {loading && (
          <div className="flex flex-row space-x-2 items-center justify-center">
            <span>Processing</span>{" "}
            <Loader2 className="animate-spin mr-2 h-4 w-4" />
          </div>
        )}
      </div>
      <CommandPalette manifests={manifests} runManifest={runManifest} />
    </React.Fragment>
  );
}

export default App;
