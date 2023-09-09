package manifest

import "strings"

type Argument string

const (
	ContentArgument Argument = "{{args.content}}"
	LinesArgument   Argument = "{{args.lines}}"
)

func (a Argument) String() string {
	return string(a)
}

func (a Argument) IsContent() bool {
	return a == ContentArgument
}

func (a Argument) IsLines() bool {
	return a == LinesArgument
}

func (a Argument) IsPlain() bool {
	return !a.IsContent() && !a.IsLines()
}

func (a Argument) Value(input *string) []string {
	if a.IsContent() {
		return []string{*input}
	}
	if a.IsLines() {
		return strings.Split(*input, "\n")
	}
	return []string{a.String()}
}
