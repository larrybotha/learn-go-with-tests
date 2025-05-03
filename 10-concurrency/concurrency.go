package concurrency

type (
	WebsiteChecker func(string) bool
	// anonymous struct
	result struct {
		string
		bool
	}
)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			// send the result to the channel
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range len(urls) {
		// when the channel receives a value, assign it to r
		r := <-resultChannel

		// update the map
		results[r.string] = r.bool
	}

	return results
}
