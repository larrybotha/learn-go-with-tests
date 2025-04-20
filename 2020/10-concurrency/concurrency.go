package concurrency

type WebsiteChecker func(string) bool

/*
one can use a struct with unnamed fields

these fields can be accessed using the type of the field as the name

e.g.

myVar := struct{int,bool}{5,true}
number := myVar.int
*/
type result struct {
	string
	bool
}

func CheckWebsitesSlow(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	/*
		create a channel using `make`

		chan expects a type describing the contents of the channel
	*/
	resultChannel := make(chan result)

	for _, url := range urls {
		/*
			`go` in front of a function call makes it a goroutine

			We use an anonymous function, and call it immediately

			We need to pass the value to the anonymous function, otherwise the only value
			that will be received is the last url in the slice. This is the same as closures
			that use `this` in Javascript
		*/
		go func(u string) {
			/*
				a channel has a value sent to it

				`<-` forms part of a 'send statement'

				a 'send statement' will send a value to a channel
			*/
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		/*
			we iterate over the length of the urls, using a 'receive expression` to assign
			the result of each value in the channel to a variable

			The value from the receive expression is of the type of the channel
		*/
		result := <-resultChannel

		results[result.string] = result.bool
	}

	return results
}
