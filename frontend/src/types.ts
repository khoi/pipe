export type ContentArgument = "{{args.content}}";
export type LinesArgument = "{{args.lines}}";

export type Argument = ContentArgument | LinesArgument;

export type Output = "clipboard+replace" | "clipboard" | "replace" | "noop";

export type Manifest = {
  name: string;
  description: string;
  pipe: {
    exec: string;
    args: Array<string | Argument>;
    stdin?: string;
  };
  output: Output;
  tags: string[];
};
