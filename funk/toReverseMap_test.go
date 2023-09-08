package funk

import (
	"reflect"
	"testing"

	"github.com/rekki/go/pkg/rekki/errors"
)

func TestToReverseMap(t *testing.T) {
	type Puppy struct {
		ID   int
		Name string
	}

	puppy1 := Puppy{ID: 1, Name: "Tess"}
	puppy2 := Puppy{ID: 2, Name: "Floss"}
	puppy3 := Puppy{ID: 3, Name: "Snoopy"}

	puppyMap := map[int]*Puppy{
		1: &puppy1,
		2: &puppy2,
		3: &puppy3,
	}

	dupePuppyMap := map[int]*Puppy{
		1: &puppy1,
		2: &puppy2,
		3: &puppy2,
	}

	tests := []struct {
		name     string
		puppyMap map[int]*Puppy
		wantMap  map[*Puppy]int
		wantErr  error
	}{
		{
			name:     "map by ID",
			puppyMap: puppyMap,
			wantMap: map[*Puppy]int{
				&puppy1: 1,
				&puppy2: 2,
				&puppy3: 3,
			},
			wantErr: nil,
		},

		{
			name:     "all empty",
			puppyMap: map[int]*Puppy{},
			wantMap:  map[*Puppy]int{},
			wantErr:  nil,
		},

		{
			name: "some empty",
			puppyMap: map[int]*Puppy{
				1: &puppy1,
				2: nil,
				3: &puppy3,
			},
			wantMap: map[*Puppy]int{
				&puppy1: 1,
				&puppy3: 3,
			},
			wantErr: nil,
		},
		{
			name:     "some with duplicate values",
			puppyMap: dupePuppyMap,
			wantMap:  map[*Puppy]int{},
			wantErr:  errors.New("Input Map has repeated values"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMap, gotErr := ToReverseMap(tt.puppyMap)
			if tt.wantErr != nil {
				if !reflect.DeepEqual(gotErr.Error(), tt.wantErr.Error()) {
					t.Errorf("Expected Err %s, Got Error %s", tt.wantErr, gotErr)
				}
			} else if !reflect.DeepEqual(gotMap, tt.wantMap) {
				t.Errorf("ToReverseMap() = %v, want %v\t[%v] len %d %d", gotMap, tt.wantMap, tt.name, len(gotMap), len(tt.wantMap))
			}
		})
	}
}

func TestToReverseArrayMap(t *testing.T) {
	type Puppy struct {
		ID   int
		Name string
	}

	puppy1 := Puppy{ID: 1, Name: "Tess"}
	puppy2 := Puppy{ID: 2, Name: "Floss"}
	puppy3 := Puppy{ID: 3, Name: "Snoopy"}

	puppyMap := map[int]*Puppy{
		1: &puppy1,
		2: &puppy2,
		3: &puppy3,
	}

	dupePuppyMap := map[int]*Puppy{
		1: &puppy1,
		2: &puppy2,
		3: &puppy2,
	}

	tests := []struct {
		name     string
		puppyMap map[int]*Puppy
		wantMap  map[*Puppy][]int
	}{
		{
			name:     "map by ID",
			puppyMap: puppyMap,
			wantMap: map[*Puppy][]int{
				&puppy1: {1},
				&puppy2: {2},
				&puppy3: {3},
			},
		},

		{
			name:     "all empty",
			puppyMap: map[int]*Puppy{},
			wantMap:  map[*Puppy][]int{},
		},

		{
			name: "some empty",
			puppyMap: map[int]*Puppy{
				1: &puppy1,
				2: nil,
				3: &puppy3,
			},
			wantMap: map[*Puppy][]int{
				&puppy1: {1},
				nil:     {2},
				&puppy3: {3},
			},
		},
		{
			name:     "some with duplicate values",
			puppyMap: dupePuppyMap,
			wantMap: map[*Puppy][]int{
				&puppy1: {1},
				&puppy2: {2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMap := ToReverseArrayMap(tt.puppyMap)
			if !reflect.DeepEqual(gotMap, tt.wantMap) {
				for key := range gotMap {
					if !IdenticalSets(tt.wantMap[key], gotMap[key]) {
						t.Errorf("ToReverseArrayMap() = %v, want %v\t[%v] len %d %d", gotMap, tt.wantMap, tt.name, len(gotMap), len(tt.wantMap))
					}
				}
			}
		})
	}
}

func TestToBespokeReverseMap(t *testing.T) {
	type Puppy struct {
		ID    int
		Name  string
		Breed func(int) int
	}

	getPuppyName := func(puppy *Puppy) string {
		if puppy == nil {
			return ""
		}
		return puppy.Name
	}

	double := func(x int) int {
		return x * 2
	}

	puppyMap := map[int]*Puppy{
		1: {ID: 1, Name: "puppy1", Breed: double},
		2: {ID: 2, Name: "puppy2", Breed: double},
		3: {ID: 3, Name: "puppy3", Breed: double},
	}

	dupePuppyMap := map[int]*Puppy{
		1: {ID: 1, Name: "puppy1", Breed: double},
		2: {ID: 2, Name: "puppy1", Breed: double},
		3: {ID: 3, Name: "puppy3", Breed: double},
	}

	tests := []struct {
		name     string
		puppyMap map[int]*Puppy
		newKey   func(*Puppy) string
		wantMap  map[string]int
		wantErr  error
	}{
		{
			name:     "reverse map by name",
			puppyMap: puppyMap,
			newKey:   getPuppyName,
			wantMap: map[string]int{
				"puppy1": 1,
				"puppy2": 2,
				"puppy3": 3,
			},
			wantErr: nil,
		},

		{
			name:     "all empty",
			puppyMap: map[int]*Puppy{},
			wantMap:  map[string]int{},
			newKey:   getPuppyName,
			wantErr:  nil,
		},

		{
			name:     "some with duplicate values",
			puppyMap: dupePuppyMap,
			wantMap:  map[string]int{},
			newKey:   getPuppyName,
			wantErr:  errors.New("Input Map has repeated values"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMap, gotErr := ToBespokeReverseMap(tt.puppyMap, tt.newKey)
			if tt.wantErr != nil {
				if !reflect.DeepEqual(gotErr.Error(), tt.wantErr.Error()) {
					t.Errorf("Expected Err %s, Got Error %s", tt.wantErr, gotErr)
				}
			} else if !reflect.DeepEqual(gotMap, tt.wantMap) {
				t.Errorf("ToReverseMap() = %v, want %v\t[%v] len %d %d", gotMap, tt.wantMap, tt.name, len(gotMap), len(tt.wantMap))
			}
		})
	}
}
