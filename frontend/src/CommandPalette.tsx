import React from "react";

import { Calendar } from "lucide-react";

import {
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import { manifest } from "wailsjs/go/models";

interface CommandPaletteProps {
  manifests: manifest.Manifest[];
  runManifest: (manifest: manifest.Manifest) => void;
}

export function CommandPalette({
  manifests,
  runManifest,
}: CommandPaletteProps) {
  const [open, setOpen] = React.useState(false);

  React.useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === "Escape") {
        setOpen(false);
      }
      if (e.key === "k" && (e.metaKey || e.ctrlKey)) {
        e.preventDefault();
        setOpen((open) => !open);
      }
    };

    document.addEventListener("keydown", down);
    return () => document.removeEventListener("keydown", down);
  }, []);

  return (
    <CommandDialog open={open} onOpenChange={setOpen}>
      <CommandInput placeholder="Type a command or search..." />
      <CommandList>
        <CommandEmpty>No results found.</CommandEmpty>
        <CommandGroup heading="Pipes">
          {manifests.map((manifest) => {
            return (
              <CommandItem
                key={manifest.id}
                onSelect={() => {
                  runManifest(manifest);
                  setOpen(false);
                }}
              >
                <Calendar className="mr-2 h-4 w-4" />
                <span>{manifest.name}</span>
              </CommandItem>
            );
          })}
        </CommandGroup>
      </CommandList>
    </CommandDialog>
  );
}
