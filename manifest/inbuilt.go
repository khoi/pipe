package manifest

import "golang.org/x/net/context"

type Inbuilt = func(ctx context.Context, input *string) (string, error)
