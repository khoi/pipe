import debounce from "lodash.debounce";
import { ModelOperations } from "@vscode/vscode-languagedetection";

import weightsURL from "./group1-shard1of1.bin?url";
import model from "./model.json";

const fetchWeights = async () => {
  const res = await fetch(weightsURL);
  const buffer = await res.arrayBuffer();
  return buffer;
};

const modulOperations = new ModelOperations({
  modelJsonLoaderFunc: () => Promise.resolve(model),
  weightsLoaderFunc: fetchWeights,
});

export const detectLanguage = debounce(async function detectLanguage(
  value: string,
  onResult: (result: string) => void,
) {
  if (!value) {
    return;
  }
  try {
    const result = await modulOperations.runModel(value);
    if (Array.isArray(result) && result.length > 0) {
      onResult(result[0].languageId);
    }
  } catch (error) {
    console.error(error);
  }
}, 2000);
