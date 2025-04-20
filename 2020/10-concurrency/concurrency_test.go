package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsitesSlow(t *testing.T) {
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

		got := CheckWebsitesSlow(mockWebsiteChecker, urls)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func TestCheckWebsites(t *testing.T) {
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
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

/*
benchmarks
*/

func BenchmarkCheckWebsitesSlow(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "foo"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsitesSlow(stubWebsiteChecker, urls)
	}
}

func stubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}
