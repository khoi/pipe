package manifest

import (
	"strconv"
	"time"

	"golang.org/x/net/context"
)

func currentUnixTimeStamp(_ context.Context, input *string) (string, error) {
	return strconv.FormatInt(time.Now().Unix(), 10), nil
}

var CurrentTimeStamp = Manifest{
	ID:          "current-unix-timestamp",
	Name:        "Current Unix Time Stamp",
	Description: "Get current Unix time stamp",
	Pipe: Pipe{
		Handler: currentUnixTimeStamp,
	},
	Output: Replace,
	Tags:   []string{"timestamp", "date"},
}
