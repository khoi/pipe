export type SelectionArgument = "{{args.selection}}";
export type LinesArgument = "{{args.lines}}";

export type Argument = SelectionArgument | LinesArgument;

export type Manifest = {
  name: string;
  description: string;
  pipe: {
    exec: string;
    args: Array<string | Argument>;
    stdin?: string;
  };
  output: "clipboard+replace" | "clipboard" | "replace" | "noop";
  tags?: string[];
};
