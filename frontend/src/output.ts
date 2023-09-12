import type { Output } from "./types";

export const write = (output: string, strategy: Output) => {
  switch (strategy) {
    case 'clipboard':
      navigator.clipboard.writeText(output)
      break
    case 'clipboard+replace':
      navigator.clipboard.writeText(output)
      return true
    case 'replace':
      return true
    default:
      return false
  }
}
