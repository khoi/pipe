package funk

import (
	"reflect"
	"testing"
)

func TestToBespokeMap(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	getUserID := func(user *User) int {
		if user == nil {
			return 0
		}
		return user.ID
	}
	getUserName := func(user *User) string {
		if user == nil {
			return ""
		}
		return user.Name
	}

	users := []*User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
		{ID: 3, Name: "Jack"},
	}

	tests := []struct {
		name  string
		args  []*User
		key   func(*User) int
		value func(*User) string
		want  map[int]string
	}{
		{
			name:  "map by ID",
			args:  users,
			key:   getUserID,
			value: getUserName,
			want: map[int]string{
				1: "John",
				2: "Jane",
				3: "Jack",
			},
		},
		{
			name:  "all empty",
			args:  []*User{},
			key:   getUserID,
			value: getUserName,
			want:  map[int]string{},
		},
		{
			name: "some empty",
			args: []*User{
				{ID: 1, Name: "John"},
				nil,
				{ID: 3, Name: "Jack"},
			},
			key:   getUserID,
			value: getUserName,
			want: map[int]string{
				0: "",
				1: "John",
				3: "Jack",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapBespoke(tt.args, tt.key, tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v\t[%v]", got, tt.want, tt.name)
			}
		})
	}
}
