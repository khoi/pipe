package funk

import (
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
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

	users := []*User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
		{ID: 3, Name: "Jack"},
	}

	tests := []struct {
		name string
		args []*User
		key  func(*User) int
		want map[int]*User
	}{
		{
			name: "map by ID",
			args: users,
			key:  getUserID,
			want: map[int]*User{
				1: {ID: 1, Name: "John"},
				2: {ID: 2, Name: "Jane"},
				3: {ID: 3, Name: "Jack"},
			},
		},
		{
			name: "all empty",
			args: []*User{},
			key:  getUserID,
			want: map[int]*User{},
		},
		{
			name: "some empty",
			args: []*User{
				{ID: 1, Name: "John"},
				nil,
				{ID: 3, Name: "Jack"},
			},
			key: getUserID,
			want: map[int]*User{
				0: nil,
				1: {ID: 1, Name: "John"},
				3: {ID: 3, Name: "Jack"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMap(tt.args, tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v\t[%v]", got, tt.want, tt.name)
			}
		})
	}
}

func TestToArrayMap(t *testing.T) {
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

	users := []*User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
		{ID: 3, Name: "Jack"},
	}

	tests := []struct {
		name string
		args []*User
		key  func(*User) int
		want map[int][]*User
	}{
		{
			name: "map by ID",
			args: users,
			key:  getUserID,
			want: map[int][]*User{
				1: {{ID: 1, Name: "John"}},
				2: {{ID: 2, Name: "Jane"}},
				3: {{ID: 3, Name: "Jack"}},
			},
		},
		{
			name: "all empty",
			args: []*User{},
			key:  getUserID,
			want: map[int][]*User{},
		},
		{
			name: "some empty",
			args: []*User{
				{ID: 1, Name: "John"},
				nil,
				{ID: 3, Name: "Jack"},
			},
			key: getUserID,
			want: map[int][]*User{
				0: {nil},
				1: {{ID: 1, Name: "John"}},
				3: {{ID: 3, Name: "Jack"}},
			},
		},
		{
			name: "some dupes",
			args: []*User{
				{ID: 1, Name: "John"},
				{ID: 1, Name: "Johnnie"},
				{ID: 3, Name: "Jack"},
			},
			key: getUserID,
			want: map[int][]*User{
				1: {{ID: 1, Name: "John"}, {ID: 1, Name: "Johnnie"}},
				3: {{ID: 3, Name: "Jack"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayMap(tt.args, tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v\t[%v]", got, tt.want, tt.name)
			}
		})
	}
}
