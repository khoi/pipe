import { ModelOperations } from "@vscode/vscode-languagedetection";

import weightsURL from "./group1-shard1of1.bin?url";
import model from "./model.json";

const fetchWeights = async () => {
  const res = await fetch(weightsURL);
  const buffer = await res.arrayBuffer();
  return buffer;
};

export const modulOperations = new ModelOperations({
  modelJsonLoaderFunc: async () => model,
  weightsLoaderFunc: fetchWeights,
});
