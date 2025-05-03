package concurrency

import (
	"testing"
	"time"
)

func slowWebsiteChecker(x string) bool {
	time.Sleep(time.Millisecond * 20)

	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range len(urls) {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for b.Loop() {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
