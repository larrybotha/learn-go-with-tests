package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"foo"},
			[]string{"foo"},
		},

		{
			"struct with two fields",
			struct {
				Name string
				City string
			}{"foo", "bar"},
			[]string{"foo", "bar"},
		},

		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"foo", 1},
			[]string{"foo"},
		},

		{
			"struct with nested struct",
			Person{"foo", Profile{1, "bar"}},
			[]string{"foo", "bar"},
		},

		{
			"struct with pointer",
			&Person{"foo", Profile{1, "bar"}},
			[]string{"foo", "bar"},
		},

		{
			"struct with slices",
			[]Person{
				{"foo", Profile{1, "bar"}},
				{"baz", Profile{2, "quux"}},
			},
			[]string{"foo", "bar", "baz", "quux"},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(x string) {
				got = append(got, x)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
