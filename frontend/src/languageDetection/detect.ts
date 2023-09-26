import debounce from "lodash.debounce";
import { ModelOperations } from "@vscode/vscode-languagedetection";

import weightsURL from "./group1-shard1of1.bin?url";
import model from "./model.json";
import { LanguageSupport, StreamLanguage } from "@codemirror/language";
import { langs } from "@uiw/codemirror-extensions-langs";

const fetchWeights = async () => {
  const res = await fetch(weightsURL);
  const buffer = await res.arrayBuffer();
  return buffer;
};

const modulOperations = new ModelOperations({
  modelJsonLoaderFunc: () => Promise.resolve(model),
  weightsLoaderFunc: fetchWeights,
});

export type DetectionResult = StreamLanguage<unknown> | LanguageSupport | null;

const langMap = {
  ts: langs.typescript(),
  js: langs.javascript({ jsx: true }),
  html: langs.html(),
  css: langs.css(),
  md: langs.markdown(),
  json: langs.json(),
};

export const detectLanguage = debounce(async function detectLanguage(
  value: string,
  onResult: (result: DetectionResult) => void
) {
  if (!value) {
    return;
  }
  try {
    const result = await modulOperations.runModel(value);
    if (Array.isArray(result) && result.length > 0) {
      // TODO: convert result[0].languageId to language from langMap
      onResult(langMap.json);
    } else {
      onResult(null);
    }
  } catch (error) {
    console.error(error);
  }
},
2000);
