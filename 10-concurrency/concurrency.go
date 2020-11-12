package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsitesSlow(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		/*
			`go` in front of a function call makes it a goroutine

			We use an anonymous function, and call it immediately

			We need to pass the value to the anonymous function, otherwise the only value
			that will be received is the last url in the slice. This is the same as closures
			that use `this` in Javascript
		*/
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}
