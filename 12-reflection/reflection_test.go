package reflection

import (
	"reflect"
	"slices"
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
		Input         any
		ExpectedCalls []string
	}{
		{
			Name:  "struct with no fields",
			Input: struct{}{},
		},
		{
			Name:          "struct with one string field",
			Input:         struct{ Name string }{"Sam"},
			ExpectedCalls: []string{"Sam"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Sam", "Jhb"},
			ExpectedCalls: []string{"Sam", "Jhb"},
		},
		{
			Name: "struct with an int",
			Input: struct {
				Name string
				Age  int
			}{"Sam", 30},
			ExpectedCalls: []string{"Sam"},
		},
		{
			Name: "struct with nested struct",
			Input: Person{
				Name: "Sam",
				Profile: Profile{
					Age:  30,
					City: "Bar",
				},
			},
			ExpectedCalls: []string{"Sam", "Bar"},
		},
		{
			Name: "struct with pointer to nested struct",
			Input: &Person{
				Name:    "Sam",
				Profile: Profile{30, "Bar"},
			},
			ExpectedCalls: []string{"Sam", "Bar"},
		},
		{
			Name: "with slices",
			Input: []Profile{
				{33, "Foo"},
				{50, "Bar"},
			},
			ExpectedCalls: []string{"Foo", "Bar"},
		},
		{
			Name: "with arrays",
			Input: [2]Profile{
				{33, "Foo"},
				{50, "Bar"},
			},
			ExpectedCalls: []string{"Foo", "Bar"},
		},
	}

	for _, x := range cases {
		t.Run(x.Name, func(t *testing.T) {
			var got []string

			walk(x.Input, func(v string) {
				got = append(got, v)
			})

			if !reflect.DeepEqual(got, x.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, x.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		value := map[string]string{"foo": "bar", "baz": "quux"}
		want := []string{"bar", "quux"}
		var got []string

		walk(value, func(x string) {
			got = append(got, x)
		})

		for _, x := range got {
			assertContains(t, want, x)
		}
	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{33, "foo"}
			ch <- Profile{50, "bar"}
			close(ch)
		}()

		var got []string
		want := []string{"foo", "bar"}

		walk(ch, func(x string) {
			got = append(got, x)
		})

		for _, x := range want {
			assertContains(t, got, x)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		value := func() (Profile, Profile) {
			return Profile{33, "foo"}, Profile{30, "bar"}
		}
		want := []string{"foo", "bar"}
		var got []string

		walk(value, func(x string) {
			got = append(got, x)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, xs []string, x string) {
	t.Helper()

	if !slices.Contains(xs, x) {
		t.Errorf("expected %q to be in %v", x, xs)
	}
}
