import debounce from "lodash.debounce";
import { ModelOperations } from "@vscode/vscode-languagedetection";

import weightsURL from "./group1-shard1of1.bin?url";
import model from "./model.json";
import { LanguageSupport, StreamLanguage } from "@codemirror/language";
import {
  langNames,
  langs,
  LanguageName,
} from "@uiw/codemirror-extensions-langs";

const fetchWeights = async () => {
  const res = await fetch(weightsURL);
  const buffer = await res.arrayBuffer();
  return buffer;
};

const modulOperations = new ModelOperations({
  modelJsonLoaderFunc: () => Promise.resolve(model),
  weightsLoaderFunc: fetchWeights,
});

function isValidLanguage(lang: string): lang is LanguageName {
  return langNames.includes(lang as LanguageName);
}

export type DetectionResult = StreamLanguage<unknown> | LanguageSupport;

export const detectLanguage = debounce(async function detectLanguage(
  value: string,
  onResult: (result: DetectionResult) => void,
) {
  if (!value) {
    return;
  }
  try {
    const result = await modulOperations.runModel(value);
    if (Array.isArray(result) && result.length > 0) {
      const lang = result[0].languageId;
      if (!isValidLanguage(lang)) return;
      const langSupport = langs[lang]();
      onResult(langSupport);
    }
  } catch (error) {
    console.error(error);
  }
}, 2000);
