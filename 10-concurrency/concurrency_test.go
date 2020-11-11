package concurrency

import (
	"reflect"
	"testing"
)

func TestWebsiteChecker(t *testing.T) {
	t.Run("checks websites", func(t *testing.T) {
		urls := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}

		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		got := CheckWebsites(mockWebsiteChecker, urls)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}
