package funk

import "testing"

func TestReplaceAllCopies(t *testing.T) {
	type args struct {
		text        string
		toReplace   string
		replaceWith string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"replace all copies for a single instance", args{"finished finickety", "fini", ""}, "shed ckety"},
		{"replace all copies for a single instance", args{"fifininished fifininickety", "fini", ""}, "shed ckety"},
		{"replace all multiple spaces with a single space", args{"a  b  c       d e", "  ", " "}, "a b c d e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAllRecursively(tt.args.text, tt.args.toReplace, tt.args.replaceWith); got != tt.want {
				t.Errorf("ReplaceAllRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}
