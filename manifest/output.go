package manifest

import "context"

type Output string

const (
	ClipboardReplace Output = "clipboard+replace"
	Clipboard        Output = "clipboard"
	Replace          Output = "replace"
	Noop             Output = "noop"
)

func (o Output) String() string {
	return string(o)
}

func (o Output) Write(ctx context.Context, output string) error {
	switch o {
	case ClipboardReplace:
		return clipboardReplace(ctx, output)
	case Clipboard:
		return copyToClipboard(ctx, output)
	case Replace:
		return replace(ctx, output)
	case Noop:
		fallthrough
	default:
		return nil
	}
}

func copyToClipboard(ctx context.Context, output string) error {
	return nil
}

func replace(ctx context.Context, output string) error {
	return nil
}

func clipboardReplace(ctx context.Context, output string) error {
	err := copyToClipboard(ctx, output)
	if err != nil {
		return err
	}
	return replace(ctx, output)
}
