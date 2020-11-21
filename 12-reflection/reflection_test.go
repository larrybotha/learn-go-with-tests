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
	t.Run("with structs, arrays, and slices", func(t *testing.T) {
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
				"slices",
				[]Profile{
					{1, "foo"},
					{2, "bar"},
				},
				[]string{"foo", "bar"},
			},

			{
				"arrays",
				[2]Profile{
					{1, "foo"},
					{2, "bar"},
				},
				[]string{"foo", "bar"},
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
	})

	t.Run("with maps", func(t *testing.T) {
		testValue := map[string]string{
			"foo": "bar",
			"baz": "quux",
		}
		var got []string
		want := []string{"bar", "quux"}

		walk(testValue, func(value string) {
			got = append(got, value)
		})

		for _, x := range want {
			assertContains(t, got, x)

		}
	})
}

func assertContains(t *testing.T, xs []string, value string) {
	t.Helper()
	isInArray := false

	for _, x := range xs {
		if x == value {
			isInArray = true
		}
	}

	if !isInArray {
		t.Errorf("expected %q, to be in %v", value, xs)
	}
}
